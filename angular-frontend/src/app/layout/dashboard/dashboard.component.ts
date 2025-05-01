import { Component, OnInit } from '@angular/core';
import { DashboardService } from '../../core/services/dashboard.service';
import { CodebasePayload, IUserChatSessionsData } from '../../core/interfaces/dashboard.interfaces';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { NgClass } from '@angular/common';
import { ToastrService } from 'ngx-toastr';
import { HeaderComponent } from '../header/header.component';
import { Router } from '@angular/router';

@Component({
  selector: 'app-dashboard',
  imports: [NgClass, ReactiveFormsModule, HeaderComponent],
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.scss'
})
export class DashboardComponent implements OnInit {
  codebases: CodebasePayload[] = [];
  userChatSessions: IUserChatSessionsData[] = [];
  selectedCodebaseIds: string[] = [];
  extractCodeBaseClicked: boolean = false;
  codebaseForm: FormGroup;
  submitted = false;
  codeBaseSyncingStarted = false;
  constructor(
    private dashBoardService: DashboardService,
    private fb: FormBuilder,
    private toastr: ToastrService,
    private router: Router,
  ) {
    this.codebaseForm = this.fb.group({
      codeBaseName: ['', Validators.required],
      gitHubURL: ['', Validators.required],
      username: ['', Validators.required],
      token: ['', Validators.required],
      branch: ['', Validators.required],
      folderPath: ['', Validators.required],
    });
  }
  ngOnInit() {
    this.CompoenentDataSetup()

  }
  get f() {
    return this.codebaseForm.controls;
  }
  private async CompoenentDataSetup() {
    await this.getCodeBases()
    await this.retrieveChatSessions();
  }
  private async getCodeBases() {
    this.codebases = await this.dashBoardService.retrieveCodeBaseData();
    console.log(`CodeBase Data : ${this.codebases}`)
  }
  private async retrieveChatSessions() {
    this.userChatSessions = await this.dashBoardService.RetrieveUserChatSessions();
    console.log("Chat Sessions :", this.userChatSessions);
  }

  ExtractCodeBase() {
    this.extractCodeBaseClicked = true;
  }
  protected async SyncCodeBases() {
    await this.dashBoardService.SynccodeBases();
    await this.getCodeBases();
  }
  onSubmit() {
    this.submitted = true;
    if (this.codebaseForm.invalid) return;

    const payload = {
      codebase_name: this.codebaseForm.value.codeBaseName,
      github_url: this.codebaseForm.value.gitHubURL,
      username: this.codebaseForm.value.username,
      token: this.codebaseForm.value.token,
      branch: this.codebaseForm.value.branch,
      folder_path: this.codebaseForm.value.folderPath
    };
    console.log('Submitting:', payload);
    this.codeBaseSyncingStarted = true;
    this.dashBoardService.extractCode(payload).subscribe({
      next: () => {
        this.toastr.success('Code Base Synced  Succesfully', 'Success');
        this.codeBaseSyncingStarted = false;
        this.extractCodeBaseClicked = false;
        this.getCodeBases();
      },
      error: (error) => {
        console.log(error);
        this.toastr.error('CodeBase Syncing Failed ...', error.error.detail);
        this.codeBaseSyncingStarted = false;
        this.extractCodeBaseClicked = false;
      },
    })
  }
  StartNewChatSession() {
    this.dashBoardService.createChatSession(this.selectedCodebaseIds).subscribe({
      next: ({ sessionId }) => {
        this.navigateToChatSession(sessionId);
      },
      error: (err) => {
        console.error('Failed to create chat session', err);
        this.toastr.error("Failed to Create Chat Session")
      }
    });
  }

  navigateToChatSession(sessionId: string) {
    this.router.navigate(['/chat-session'], {
      queryParams: {
        sessionId: sessionId,
      }
    });
  }

  deleteCodeBaseContext(codeBaseId: string) {
    this.dashBoardService.DeleteCodeBasdeContext(codeBaseId).subscribe({
      next: async () => {
        await this.getCodeBases()
        this.toastr.success(`Code Base Deleted entirely  ...`)
      },
      error: (error) => {
        console.log(error);
        this.toastr.error('Error in deleting codeBase .');
      }
    })
  }

  deleteChatSession(sessionId: string) {
    this.dashBoardService.DeleteChatSession(sessionId).subscribe({
      next: async () => {
        await this.retrieveChatSessions()
        this.toastr.success(`Chat Session Deleted`)
      },
      error: (error) => {
        console.log(error);
        this.toastr.error('Error in deleting Chat Session');
      }
    })
  }

  toggleCodebaseSelection(codebaseId: string): void {
    const index = this.selectedCodebaseIds.indexOf(codebaseId);
    if (index !== -1) {
      this.selectedCodebaseIds.splice(index, 1);
    } else {
      this.selectedCodebaseIds.push(codebaseId);
    }
    console.log('Selected codebase IDs:', this.selectedCodebaseIds);
  }
  toggleAllCodebases(event: any): void {
    if (event.target.checked) {
      this.selectedCodebaseIds = this.codebases.map(codebase => codebase.codebase_id);
    } else {
      this.selectedCodebaseIds = [];
    }
    console.log('Selected codebase IDs:', this.selectedCodebaseIds);
  }
  isCodebaseSelected(codebaseId: string): boolean {
    return this.selectedCodebaseIds.includes(codebaseId);
  }

  closeCodeBaseClicked() {
    this.extractCodeBaseClicked = false;
  }

}

import { Component, Input } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ReferencesWithSimilarity } from '../../../core/interfaces/chat.service.interface';
import { ToastrService } from 'ngx-toastr';
import { MarkdownComponent } from 'ngx-markdown';

@Component({
  selector: 'app-code-modal',
  imports: [MarkdownComponent],
  templateUrl: './code-modal.component.html',
  styleUrl: './code-modal.component.scss'
})
export class CodeModalComponent {
  constructor(private modalService: NgbModal, private toastr: ToastrService) { }
  @Input() reference  !: ReferencesWithSimilarity;
  formatBody(text: string): string {
    return text.replace(/\n/g, '<br>');
  }
  copyCode(): void {
    // Get the raw code text without HTML formatting
    const codeText = this.reference.code;

    // Use the Clipboard API to copy the text
    navigator.clipboard.writeText(codeText)
      .then(() => {
        this.toastr.success('Succesfully Copied  .', 'Success');
      })
      .catch(err => {
        console.error('Failed to copy text: ', err);
      });
  }
  closeModal() {
    this.modalService.dismissAll();
  }
}

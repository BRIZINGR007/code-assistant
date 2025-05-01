import { Component, EventEmitter, Input, Output } from '@angular/core';
import { ReferencesWithSimilarity } from '../../core/interfaces/chat.service.interface';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { CodeModalComponent } from '../../shared/modals/code-modal/code-modal.component';

@Component({
  selector: 'app-refrence-traversal',
  imports: [],
  templateUrl: './refrence-traversal.component.html',
  styleUrl: './refrence-traversal.component.scss'
})
export class RefrenceTraversalComponent {
  @Input() referencesForTraversal: ReferencesWithSimilarity[] = [];
  @Output() CloseReferenceTraversalEmitter = new EventEmitter<any>();

  constructor(private modalService: NgbModal) { }

  CloseRefernces() {
    this.CloseReferenceTraversalEmitter.emit();
  }
  toPercentage(value: number) {
    const percentage = (value * 100).toFixed(0);
    return `${percentage}%`;
  }
  OpenCodeView(index: number) {
    const modalRef = this.modalService.open(CodeModalComponent, { size: 'xl', centered: true });
    modalRef.componentInstance.reference = this.referencesForTraversal[index];
  }


}

import { CommonModule, NgComponentOutlet } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';

@Component({
  selector: 'app-modal',
  standalone: true,
  imports: [NgComponentOutlet, CommonModule],
  templateUrl: './modal.component.html',
  styleUrl: './modal.component.css',
})
export class ModalComponent {
  @Input() modalTitle: string = '';

  @Input() isVisible: boolean = false;

  @Output() closeModal = new EventEmitter<any>();

  close() {
    this.closeModal.emit();
  }
}

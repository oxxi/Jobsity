import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';
import { NgIconComponent, provideIcons } from '@ng-icons/core';
import { heroTrash, heroPencilSquare } from '@ng-icons/heroicons/outline';
@Component({
  selector: 'app-todo-item',
  standalone: true,
  imports: [NgIconComponent, CommonModule],
  viewProviders: [provideIcons({ heroTrash, heroPencilSquare })],
  templateUrl: './todo-item.component.html',
  styleUrl: './todo-item.component.css',
})
export class TodoItemComponent {
  @Input() data: any;

  @Output() delete = new EventEmitter<any>();
  @Output() edit = new EventEmitter<any>();
  @Output() complete = new EventEmitter<any>();
  @Output() inComplete = new EventEmitter<any>();

  onDelete(item: any) {
    this.delete.emit(item);
  }

  onEdit(item: any) {
    this.edit.emit(item);
  }

  onComplete(item: any) {
    this.complete.emit(item);
  }

  onIncomplete(item: any) {
    this.inComplete.emit(item);
  }
}

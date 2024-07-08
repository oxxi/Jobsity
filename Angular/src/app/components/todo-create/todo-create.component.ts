import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';
import {
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';

@Component({
  selector: 'app-todo-create',
  standalone: true,
  imports: [ReactiveFormsModule, CommonModule],

  templateUrl: './todo-create.component.html',
  styleUrl: './todo-create.component.css',
})
export class TodoCreateComponent {
  @Input() isEdit: boolean = false;
  @Input() taskEdit: any = {};
  @Output() add = new EventEmitter<any>();
  @Output() edit = new EventEmitter<any>();

  formGroup = new FormGroup({
    title: new FormControl('', {
      validators: [
        Validators.pattern('^[a-zA-Z0-9@.?_\\-\\s]+$'),
        Validators.required,
      ],
    }),
    status: new FormControl(false),
  });

  onAddTask(): void {
    if (this.formGroup.valid) {
      const title = this.formGroup.controls.title.value;
      const status = this.formGroup.controls.status.value;
      this.add.emit({ title, status });
      this.formGroup.reset();
    }
  }

  onEditTask(): void {
    if (this.formGroup.valid) {
      const title = this.formGroup.controls.title.value;

      this.edit.emit({
        id: this.taskEdit.id,
        title,
        status: this.taskEdit.status,
      });
      this.formGroup.reset();
    }
  }
}

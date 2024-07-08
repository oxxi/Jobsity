import { Component, inject, OnInit } from '@angular/core';

import { TodoItemComponent } from '../../components/todo-item/todo-item.component';
import { TodoCreateComponent } from '../../components/todo-create/todo-create.component';
import { ModalComponent } from '../../components/modal/modal.component';
import { NgIconComponent, provideIcons } from '@ng-icons/core';
import { heroPlus } from '@ng-icons/heroicons/outline';
import { TodoApiService } from '../../services/todo-api.service';
import { catchError, map, Observable, of, throwError } from 'rxjs';
import { AsyncPipe, JsonPipe } from '@angular/common';
import { TaskModel } from '../../models/taskmodel';
import { NotExpr } from '@angular/compiler';
@Component({
  selector: 'app-todo-page',
  standalone: true,
  imports: [
    TodoItemComponent,
    TodoCreateComponent,
    ModalComponent,
    AsyncPipe,
    JsonPipe,
    NgIconComponent,
  ],
  viewProviders: [provideIcons({ heroPlus })],
  templateUrl: './todo-page.component.html',
  styleUrl: './todo-page.component.css',
})
export class TodoPageComponent implements OnInit {
  dataSource$: Observable<TaskModel[]> | undefined;
  openModal: boolean = false;
  modalTitle: string = 'Add Task';
  isEditable: boolean = false;
  taskEdit: any = {};

  private service = inject(TodoApiService);

  ngOnInit(): void {
    this.init();
  }

  showModal() {
    this.modalTitle = 'Add Task';
    this.openModal = true;
  }

  onAddTask(item: any) {
    this.service.save(item).subscribe({
      next: (response) => {
        this.init();
      },
    });

    this.openModal = false;
    this.isEditable = false;
  }

  onEditTask(item: any) {
    console.log({ item });
    this.service.update(item.id, item).subscribe({
      next: (response) => {
        this.init();
      },
    });
    this.openModal = false;
    this.isEditable = false;
  }

  onCompleteTask(item: any) {
    item.status = true;
    this.service.update(item.id, item).subscribe({
      next: (response) => {
        this.init();
      },
      error(err) {
        console.error(err);
      },
    });
  }

  onIncompleteTask(item: any) {
    item.status = false;
    this.service.update(item.id, item).subscribe({
      next: (response) => {
        this.init();
      },
      error(err) {
        console.error(err);
      },
    });
  }

  onOpenEdit(item: any) {
    this.modalTitle = 'Edit Task';
    this.isEditable = true;
    this.taskEdit = item;
    this.openModal = true;
  }

  onDeleteTask(item: any) {
    this.service.delete(item.id).subscribe({
      next: (response) => {
        this.init();
      },
      error(err) {
        console.error(err);
      },
    });
  }

  onCloseModal() {
    this.openModal = false;
  }

  private init(): void {
    this.dataSource$ = this.service.getAll().pipe(
      catchError((err) => {
        console.error(err);
        return of([]);
      })
    );
  }
}

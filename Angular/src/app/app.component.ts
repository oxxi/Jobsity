import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { TodoPageComponent } from './pages/todo-page/todo-page.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, TodoPageComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
})
export class AppComponent {
  title = 'jobsity-app';
}

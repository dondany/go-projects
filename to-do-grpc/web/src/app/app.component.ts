import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet],
  template: `
    <div class="w-full flex flex-col items-center">
      <div class="w-96 flex flex-col items-center">
        <h1 class="text-4xl">My To Do App</h1>
        <div class="w-full">
          <router-outlet />
        </div>
      </div>
    </div>
  `,
  styles: [],
})
export class AppComponent {
  title = 'web';
}

import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  standalone: true,
  selector: 'app-home',
  template: `
  <div>
    <span class="bg-red-500">Home</span>
    <router-outlet />
  </div>
  `,
  imports: [CommonModule, RouterOutlet],
})
export default class ListsComponent  {
}

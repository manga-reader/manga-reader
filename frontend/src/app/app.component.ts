import { Component, OnInit } from '@angular/core';
import { MenuItem, PrimeNGConfig } from 'primeng/api';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  items: MenuItem[] = [];
  search = '';

  constructor(private primengConfig: PrimeNGConfig) {}

  ngOnInit() {
      this.primengConfig.ripple = true;
      this.items = [
        {
            label: 'My Favorite',
            icon: 'pi pi-bookmark',
        },
        {
            label: 'Latest Update',
            icon: 'pi pi-arrow-circle-up',
        }
    ];
  }
}

import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { Router } from '@angular/router';
import { MenuItem } from 'primeng/api';
import { MenubarEnum } from '../shared/menubar.enum';
import { MenubarOption } from '../shared/models/menubar-option.model';

@Component({
  selector: 'app-menubar',
  templateUrl: './menubar.component.html',
  styleUrls: ['./menubar.component.css']
})
export class MenubarComponent implements OnInit {

  @Output() clickMenubarEvent = new EventEmitter<MenubarOption>();
  items: MenuItem[] = [];
  keyword = '';

  constructor() { }

  ngOnInit(): void {
    this.items = [
      {
          label: 'My Favorite',
          icon: 'pi pi-bookmark',
          command: () => {
            this.clickMenubarEvent.emit({ menubarEnum: MenubarEnum.MyFavorite });
          },
      },
      {
          label: 'Latest Update',
          icon: 'pi pi-arrow-circle-up',
          command: async () => {
            this.clickMenubarEvent.emit({ menubarEnum: MenubarEnum.LatestUpdate });
          },
      }
    ];
  }

  async searchClick() {
    this.clickMenubarEvent.emit({ menubarEnum: MenubarEnum.Search, data: this.keyword });
  }
}

import { Component, OnInit, Output } from '@angular/core';
import { Router } from '@angular/router';
import { MenuItem } from 'primeng/api';
import { MangaList } from '../manga/shared/models/manga-list.model';
import { MangaService } from '../manga/shared/services/manga.service';

@Component({
  selector: 'app-menubar',
  templateUrl: './menubar.component.html',
  styleUrls: ['./menubar.component.css']
})
export class MenubarComponent implements OnInit {

  @Output() mangaList: MangaList;
  items: MenuItem[] = [];
  keyword = '';
  currentPage = 1;

  constructor(
    private router: Router,
    private mangaService: MangaService
  ) {
    this.mangaList = {
      manga: [],
      pager: [],
    }
  }

  ngOnInit(): void {
    this.items = [
      {
          label: 'My Favorite',
          icon: 'pi pi-bookmark',
          command: () => {
          },
      },
      {
          label: 'Latest Update',
          icon: 'pi pi-arrow-circle-up',
          command: async () => {
            this.mangaList = await this.mangaService.getLatestUpdate(this.currentPage);
            this.router.navigate(['/list'], {state: {'mangaList': this.mangaList}});
            this.currentPage = 1;
          },
      }
    ];
  }

  async searchClick() {
    this.mangaList = await this.mangaService.search(this.keyword);
    this.router.navigate(['/list'], {state: {'mangaList': this.mangaList}});
    this.currentPage = 1;
  }

  async changePage(page: any) {
    this.currentPage = page;
    this.mangaList = await this.mangaService.getLatestUpdate(this.currentPage);
  }
}

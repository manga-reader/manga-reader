import { Component, OnInit } from '@angular/core';
import { MenuItem, PrimeNGConfig } from 'primeng/api';
import { MangaList } from './shared/models/manga-list.model';
import { MangaService } from './shared/services/manga.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  items: MenuItem[] = [];
  mangaList: MangaList;
  keyword = '';
  currentPage = 1;

  constructor(
    private mangaService: MangaService,
    private primengConfig: PrimeNGConfig
  ) {
    this.mangaList = {
      manga: [],
      pager: [],
    }
  }

  ngOnInit() {
      this.primengConfig.ripple = true;
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
              this.currentPage = 1;
            },
        }
    ];
  }

  async searchClick() {
    this.mangaList = await this.mangaService.search(this.keyword);
    this.currentPage = 1;
  }

  changePage(page: any) {
    this.currentPage = page;
  }
}

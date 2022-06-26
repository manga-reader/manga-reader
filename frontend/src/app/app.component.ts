import { Component, OnInit } from '@angular/core';
import { MenuItem, PrimeNGConfig } from 'primeng/api';
import { Manga } from './shared/models/manga.model';
import { MangaService } from './shared/services/manga.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  items: MenuItem[] = [];
  mangaList: Manga[] = [];
  keyword = '';

  constructor(
    private mangaService: MangaService,
    private primengConfig: PrimeNGConfig
  ) {}

  ngOnInit() {
      this.primengConfig.ripple = true;
      this.items = [
        {
            label: 'My Favorite',
            icon: 'pi pi-bookmark',
            command: () => {
              this.mangaList = [];
            },
        },
        {
            label: 'Latest Update',
            icon: 'pi pi-arrow-circle-up',
            command: async () => {
              this.mangaList = await this.mangaService.getLatestUpdate(1);
            },
        }
    ];
  }

  async searchClick() {
    this.mangaList = await this.mangaService.search(this.keyword);
  }
}

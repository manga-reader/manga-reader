import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';
import { MenubarEnum } from 'src/app/shared/menubar.enum';
import { MenubarOption } from 'src/app/shared/models/menubar-option.model';
import { MangaList } from '../shared/models/manga-list.model';
import { MangaService } from '../shared/services/manga.service';

@Component({
  selector: 'app-manga-list',
  templateUrl: './manga-list.component.html',
  styleUrls: ['./manga-list.component.css']
})
export class MangaListComponent implements OnInit {

  mangaList: MangaList;
  currentPage = 1;
  menubar!: MenubarOption;
  @Output() changePageEvent = new EventEmitter<string>();

  constructor(
    private router: Router,
    private mangaService: MangaService
  ) {
    this.router.events.subscribe((event) => {
      if (event instanceof NavigationEnd) {
        const state = this.router?.getCurrentNavigation()?.extras?.state;

        if (state) {
          this.menubar = state['menubar'];
          switch (this.menubar.menubarEnum) {
            case MenubarEnum.MyFavorite:
              this.getMyFavorite();
              break;
            case MenubarEnum.LatestUpdate:
              this.getLatestUpdate();
              break;
            case MenubarEnum.Search:
              this.getSearchResult(this.menubar.data)
              break;
            default:
              throw (new Error('Wrong menuEnum'))
          }
        }
      }
    });

    this.mangaList = {
      manga: [],
      pager: [],
    }
  }

  ngOnInit(): void {

  }

  async changePage(page: string): Promise<void> {
    this.mangaList = await this.mangaService.changePage(this.menubar.menubarEnum, this.mangaList.pager, page)
  }

  getMyFavorite() {
  }

  async getLatestUpdate() {
    this.mangaList = await this.mangaService.getLatestUpdate(this.currentPage);
    this.currentPage = 1;
  }

  async getSearchResult(keyword: string) {
    this.mangaList = await this.mangaService.search(keyword);
    this.currentPage = 1;
  }
}

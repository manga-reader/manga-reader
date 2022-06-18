import { MangaImageParserService } from './shared/manga-image-parser.service';
import { Component, OnInit } from '@angular/core';
import { MangaService } from '../shared/services/manga.service';

@Component({
  selector: 'app-manga-viewer',
  templateUrl: './manga-viewer.component.html',
  styleUrls: ['./manga-viewer.component.css']
})
export class MangaViewerComponent implements OnInit {

  currentComic: any[] = [];
  page: any = 0;
  vol = 0;
  pages: any[] = [];
  currentImg = "";

  constructor(
    private mangaImageParserService: MangaImageParserService,
    private mangaService: MangaService
  ) { }

  async ngOnInit(): Promise<void> {
    let comicId = "10406";    // change comid id here
    this.currentComic = await this.mangaImageParserService.getHtmlImg(comicId, "1-1");
    this.currentImg = this.currentComic[0].Urls[0];
    this.getPages(this.currentComic[0].Urls.length);
    this.mangaService.search("咒術");
  }

  previousPage() {
    if (this.page > 0) {
      this.page -= 1;
      this.jumpPage();
    }
  }

  nextPage() {
    if (this.page < this.currentComic[this.vol].Urls.length - 1) {
      this.page += 1;
      this.jumpPage();
    }
  }

  previousVol() {
    if (this.vol > 0) {
      this.vol -= 1;
      this.page = 0;
      this.jumpPage();
      this.getPages(this.currentComic[this.vol].Urls.length);
    }
  }

  nextVol() {
    if (this.page < this.currentComic.length - 1) {
      this.vol += 1;
      this.page = 0;
      this.jumpPage()
      this.getPages(this.currentComic[this.vol].Urls.length);
    }
  }

  getPages(lastPage: number) {
    this.pages = [];
    for (let i = 0; i < lastPage; i++) {
      this.pages.push({
        label: i + "/" + (this.currentComic[this.vol].Urls.length - 1),
        value: i
      })
    }
  }

  jumpPage() {
    this.currentImg = this.currentComic[this.vol].Urls[this.page];
  }
}

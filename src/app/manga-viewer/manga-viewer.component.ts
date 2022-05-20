import { MangaPictureParserService } from './manga-picture-parser.service';
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

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
  currentPic = "";

  constructor(
    private mangaPictureParserService: MangaPictureParserService
  ) { }

  async ngOnInit(): Promise<void> {
    this.currentComic = await this.mangaPictureParserService.getHtmlPic("10406", "1-1")
    this.currentPic = this.currentComic[0].Urls[0];
    this.getPages(this.currentComic[0].Urls.length);
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
    this.currentPic = this.currentComic[this.vol].Urls[this.page];
  }
}

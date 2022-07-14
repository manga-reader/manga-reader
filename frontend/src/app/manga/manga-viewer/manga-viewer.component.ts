import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { MangaImageParserService } from './shared/manga-image-parser.service';

@Component({
  selector: 'app-manga-viewer',
  templateUrl: './manga-viewer.component.html',
  styleUrls: ['./manga-viewer.component.css']
})
export class MangaViewerComponent implements OnInit {

  currentComic: any[] = [];
  comicId!: string;
  vol!: number;
  page!: number;
  pages: any[] = [];
  currentVol = [];
  currentImg = "";

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private mangaImageParserService: MangaImageParserService
  ) { }

  async ngOnInit(): Promise<void> {
    this.comicId = this.route.snapshot.paramMap.get('id')!;    // change comic id here
    this.currentComic = await this.mangaImageParserService.getHtmlImg(this.comicId, '1-1');
    this.vol = Number(this.route.snapshot.paramMap.get('vol')!);
    this.page = Number(this.route.snapshot.paramMap.get('page')!);
    this.updateVolDetail();
  }

  previousPage() {
    if (this.page > 0) {
      this.page -= 1;
      this.jumpPage();
    }
  }

  nextPage() {
    if (this.page < this.currentVol.length - 1) {
      this.page += 1;
      this.jumpPage();
    }
  }

  previousVol() {
    if (this.vol > 0) {
      this.vol -= 1;
      this.page = 0;
      this.updateVolDetail();
      this.router.navigate(['/viewer', this.comicId, this.vol])
    }
  }

  nextVol() {
    if (this.page < this.currentComic.length - 1) {
      this.vol += 1;
      this.page = 0;
      this.updateVolDetail();
      this.router.navigate(['/viewer', this.comicId, this.vol])
    }
  }

  getPages(lastPage: number) {
    this.pages = [];
    for (let i = 0; i < lastPage; i++) {
      this.pages.push({
        label: i + "/" + (lastPage - 1),
        value: i
      })
    }
  }

  updateVolDetail() {
    this.currentVol = this.currentComic.find(x => x.vol === this.vol)?.Urls;
    this.jumpPage()
    this.getPages(this.currentVol.length);
  }

  jumpPage() {
    this.currentImg = this.currentVol[this.page];
  }
}

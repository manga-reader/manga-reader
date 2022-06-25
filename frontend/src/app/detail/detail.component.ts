import { Component, OnInit } from '@angular/core';
import { MangaDetail } from '../shared/models/manga-detail.model';
import { MangaService } from '../shared/services/manga.service';

@Component({
  selector: 'app-detail',
  templateUrl: './detail.component.html',
  styleUrls: ['./detail.component.css']
})
export class DetailComponent implements OnInit {

  comicId = '7340';
  comicPic = `https://www.comicabc.com/pics/0/${this.comicId}.jpg`
  mangaDetail!: MangaDetail;

  constructor(
    private mangaService: MangaService
  ) { }

  async ngOnInit(): Promise<void> {
    this.mangaDetail = await this.mangaService.getMangaDetail(this.comicId);
  }

  reverse() {
    this.mangaDetail.vols = this.mangaDetail.vols.reverse();
  }

}

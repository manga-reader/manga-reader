import { MangaService } from './../shared/services/manga.service';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-my-favorite',
  templateUrl: './my-favorite.component.html',
  styleUrls: ['./my-favorite.component.css']
})
export class MyFavoriteComponent implements OnInit {

  constructor(
    private mangaService: MangaService
  ) { }

  ngOnInit(): void {
    this.mangaService.search("咒術");
  }

}

import { Component, Input, OnInit } from '@angular/core';
import { Manga } from '../shared/models/manga.model';

@Component({
  selector: 'app-manga-list',
  templateUrl: './manga-list.component.html',
  styleUrls: ['./manga-list.component.css']
})
export class MangaListComponent implements OnInit {

  @Input() mangaList: Manga[] = [];

  constructor() { }

  ngOnInit(): void {
  }

}

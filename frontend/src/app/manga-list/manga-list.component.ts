import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MangaList } from '../shared/models/manga-list.model';
import { Manga } from '../shared/models/manga.model';
import { Pager } from '../shared/models/pager.model';

@Component({
  selector: 'app-manga-list',
  templateUrl: './manga-list.component.html',
  styleUrls: ['./manga-list.component.css']
})
export class MangaListComponent implements OnInit {

  @Input() mangaList: MangaList;
  @Output() changePageEvent = new EventEmitter<string>();

  constructor() {
    this.mangaList = {
      manga: [],
      pager: [],
    }
  }

  ngOnInit(): void {
  }

  changePage(page: string): void {
    this.changePageEvent.emit(page);
  }
}

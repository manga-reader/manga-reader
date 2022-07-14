import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Router } from '@angular/router';
import { MangaList } from '../shared/models/manga-list.model';

@Component({
  selector: 'app-manga-list',
  templateUrl: './manga-list.component.html',
  styleUrls: ['./manga-list.component.css']
})
export class MangaListComponent implements OnInit {

  @Input() mangaList: MangaList;
  @Output() changePageEvent = new EventEmitter<string>();

  constructor(
    private router: Router
  ) {
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

import { AfterViewChecked, AfterViewInit, Component, EventEmitter, Input, OnChanges, OnInit, Output, SimpleChanges } from '@angular/core';
import { Router } from '@angular/router';
import { MangaList } from '../shared/models/manga-list.model';

@Component({
  selector: 'app-manga-list',
  templateUrl: './manga-list.component.html',
  styleUrls: ['./manga-list.component.css']
})
export class MangaListComponent implements OnInit, OnChanges, AfterViewInit {

  mangaList: MangaList;
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
    console.log(history.state);
    if (history.state?.mangaList !== undefined) {
      this.mangaList = history.state?.mangaList;
    }
  }

  ngOnChanges(changes: SimpleChanges): void {
    console.log(history.state);
  }

  ngAfterViewInit(): void {
    console.log(history.state);
  }

  changePage(page: string): void {
    this.changePageEvent.emit(page);
  }
}

import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, ParamMap } from '@angular/router';
import { Observable, switchMap } from 'rxjs';
import { MangaDetail } from '../shared/models/manga-detail.model';
import { Vol } from '../shared/models/vol.model';
import { MangaService } from '../shared/services/manga.service';

@Component({
  selector: 'app-detail',
  templateUrl: './detail.component.html',
  styleUrls: ['./detail.component.css']
})
export class DetailComponent implements OnInit {

  comicId!: string;
  vols!: Vol[];
  mangaDetail$!: Observable<MangaDetail>;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private mangaService: MangaService
  ) { }

  async ngOnInit(): Promise<void> {
    this.mangaDetail$ = this.route.paramMap.pipe(
      switchMap((params: ParamMap) => {
        this.comicId = params.get('id')!;
        const detail = this.mangaService.getMangaDetail(this.comicId)
        detail.then(x =>
          this.vols = x.vols
        );
        return detail;
      })
    );
  }

  reverse() {
    this.vols = this.vols.reverse();
  }

  gotoViewer(vol: string) {
    window.open(`/viewer/${this.comicId}/${vol}`);
  }

}

import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MenuItem, PrimeNGConfig } from 'primeng/api';
import { MangaList } from './manga/shared/models/manga-list.model';
import { MangaService } from './manga/shared/services/manga.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  constructor(
    private primengConfig: PrimeNGConfig,
    private router: Router
  ) { }

  ngOnInit() {
    this.primengConfig.ripple = true;
  }

  menubarClick(event: any) {
    this.router.navigate(['/list'], { state:{menubar: event} })
  }
}

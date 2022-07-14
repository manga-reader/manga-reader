import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { MangaRoutingModule } from './manga-routing.module';
import { MangaDetailComponent } from './manga-detail/manga-detail.component';
import { CardModule } from 'primeng/card';
import { ButtonModule } from 'primeng/button';
import { MangaListComponent } from './manga-list/manga-list.component';


@NgModule({
  declarations: [
    MangaDetailComponent,
    MangaListComponent,
  ],
  imports: [
    CommonModule,
    MangaRoutingModule,
    CardModule,
    ButtonModule,
  ]
})
export class MangaModule { }

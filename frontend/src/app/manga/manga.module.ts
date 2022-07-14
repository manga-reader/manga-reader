import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { MangaRoutingModule } from './manga-routing.module';
import { MangaDetailComponent } from './manga-detail/manga-detail.component';
import { MangaListComponent } from './manga-list/manga-list.component';
import { MangaViewerComponent } from './manga-viewer/manga-viewer.component';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { DropdownModule } from 'primeng/dropdown';


@NgModule({
  declarations: [
    MangaDetailComponent,
    MangaListComponent,
    MangaViewerComponent,
  ],
  imports: [
    CommonModule,
    FormsModule,
    MangaRoutingModule,
    ButtonModule,
    CardModule,
    DropdownModule,
  ]
})
export class MangaModule { }

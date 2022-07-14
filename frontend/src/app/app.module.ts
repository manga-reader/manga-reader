import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { MangaViewerComponent } from './manga-viewer/manga-viewer.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ButtonModule } from 'primeng/button'
import { CardModule } from 'primeng/card';
import { DropdownModule } from 'primeng/dropdown';
import { FormsModule } from '@angular/forms';
import { MyFavoriteComponent } from './my-favorite/my-favorite.component';
import { DetailComponent } from './detail/detail.component';
import { MenubarModule } from 'primeng/menubar';
import { InputTextModule } from 'primeng/inputtext';
import { MangaListComponent } from './manga/manga-list/manga-list.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { MangaModule } from './manga/manga.module';

@NgModule({
  declarations: [
    AppComponent,
    MangaViewerComponent,
    MyFavoriteComponent,
    DetailComponent,
    MangaListComponent,
    PageNotFoundComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    BrowserAnimationsModule,
    CardModule,
    ButtonModule,
    DropdownModule,
    MenubarModule,
    InputTextModule,
    MangaModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

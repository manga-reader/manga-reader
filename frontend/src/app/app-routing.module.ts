import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DetailComponent } from './detail/detail.component';
import { MangaListComponent } from './manga/manga-list/manga-list.component';
import { MangaViewerComponent } from './manga-viewer/manga-viewer.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';

const routes: Routes = [
  { path: 'list', component: MangaListComponent },
  { path: 'detail/:id', component: DetailComponent },
  {
    path: 'viewer/:id/:vol',
    component: MangaViewerComponent
  },
  { path: '', redirectTo: '/list', pathMatch: 'full' },
  { path: '**', component: PageNotFoundComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

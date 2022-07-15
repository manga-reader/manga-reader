import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { MangaDetailComponent } from './manga-detail/manga-detail.component';
import { MangaListComponent } from './manga-list/manga-list.component';
import { MangaViewerComponent } from './manga-viewer/manga-viewer.component';

const routes: Routes = [
  { path: 'list', component: MangaListComponent },
  { path: 'detail/:id', component: MangaDetailComponent },
  { path: 'viewer/:id/:vol', component: MangaViewerComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class MangaRoutingModule { }

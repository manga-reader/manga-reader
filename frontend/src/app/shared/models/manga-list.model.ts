import { Manga } from "./manga.model";
import { Pager } from "./pager.model";

export interface MangaList {
  manga: Manga[],
  pager: Pager[],
}

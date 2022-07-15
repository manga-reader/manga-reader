import { Vol } from "./vol.model";

export interface MangaDetail {
  title: string,
  comicId: string,
  author: string,
  status: string,
  updateDt: string,
  description: string,
  vols: Vol[],
}

import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { firstValueFrom } from 'rxjs';
import { Manga } from '../models/manga.model';

@Injectable({
  providedIn: 'root'
})
export class MangaService {

  constructor(
    private http: HttpClient
  ) { }

  async getHtml(url: string): Promise<string> {
    return firstValueFrom(this.http.get(url, {responseType: "text"}));
  }

  search(keyword: string): Manga[] {
    const url = "http://localhost:4200/search.aspx"
    const params = new HttpParams();
    params.append("key", keyword);
    const html = firstValueFrom(this.http.get(url, {params, responseType: "text"}));
    console.log(html);
    return [];
  }

}

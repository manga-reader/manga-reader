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

  async search(keyword: string): Promise<Manga[]> {
    const url = "http://localhost:4200/search.aspx"
    let params = new HttpParams();
    params = params.append("key", keyword);
    const html = await firstValueFrom(this.http.get(url, {params, responseType: "text"}));
    console.log(html);
    const parse = this.parseManga(html);
    console.log(parse);
    return [];
  }

  parseManga(html: string): Manga[] {
    const result = html.split("<div class=\"cat2_list text-center mb-4\">")
    result.shift();
    const manga = result.map(x => {
      let title = x.split("<span>")[1].split("</span>")[0];
      title = title.replace("<font color=red>", "").replace("</font>", "");
      const comicId = x.split("<img src=\"/pics/0/")[1].split(".jpg")[0];
      const updateStatus = x.split("<span class=\"badge badge-warning font_normal\">")[1].split("</span>")[0];
      const updateDt = x.split("<li class=\"cat2_list_date\">\n")[1].split("\n</li>")[0];
      return {
        title,
        comicId,
        updateStatus,
        updateDt
      }
    })

    return manga;
  }
}

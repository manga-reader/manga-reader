import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { firstValueFrom } from 'rxjs';
import { MangaDetail } from '../models/manga-detail.model';
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

  async getMangaDetail(comicId: string): Promise<MangaDetail> {
    const html = await this.getHtml(`http://localhost:4200/html/${comicId}.html`);
    const topContent = html.split('<div class="item-top-content">')[1];
    const title = topContent.split('<h3 class="item_name" title="')[1].split('">')[0];
    const author = topContent.split('<li class="item_detail_color_gray">')[1].split('</a>\n')[1].split('\n</li>')[0];
    const status = topContent.split('href="#Comic">')[1].split('\n</div>')[0].replace('</a>', '');
    const updateDt = topContent.split('更新：<span class="font_small">')[1].split('</span>')[0];
    const description = topContent.split('<li class="item_info_detail">\n')[1].split('\n<span class="gradient"></span>')[0];
    let vols = topContent.split('<td colspan="10"></td>')[1].split('</table>')[0].split('onclick="cview(\'').slice(1).map(x => {
        const vol = x.split('.html')[0].split('-')[1];
        const name = x.split('class="Ch">\n')[1].split('</a>')[0].split('</font>')[0].replace('<font color=red id=lch>', '');

        return { vol, name }
      });
    const mangaDetail: MangaDetail = {title, comicId, author, status, updateDt, description, vols}

    return mangaDetail;
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

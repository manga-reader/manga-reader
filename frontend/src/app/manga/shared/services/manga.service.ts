import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { firstValueFrom } from 'rxjs';
import { MenubarEnum } from 'src/app/shared/menubar.enum';
import { MangaDetail } from '../models/manga-detail.model';
import { MangaList } from '../models/manga-list.model';
import { Pager } from '../models/pager.model';

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

  async getLatestUpdate(page: number): Promise<MangaList> {
    const url = `http://localhost:4200/comic/u-${page}.html`;
    const html = await this.getHtml(url);
    return this.parseMangaList(html);
  }

  async search(keyword: string): Promise<MangaList> {
    const url = "http://localhost:4200/search.aspx"
    let params = new HttpParams();
    params = params.append("key", keyword);
    const html = await firstValueFrom(this.http.get(url, {params, responseType: "text"}));
    return this.parseMangaList(html)
  }

  async changePage(menubarEnum: MenubarEnum, pager: Pager[], page: string): Promise<MangaList> {
    switch(menubarEnum) {
      case MenubarEnum.LatestUpdate:
        pager = pager.map(x => {
          let pager = x;
          pager.url = `/comic/${pager.url}`;
          return pager;
        });
        break;
      case MenubarEnum.Search:
        pager = pager.map(x => {
          let pager = x;
          pager.url = pager.url.replace('/member', '');
          return pager;
        });
        break;
    }
    const url = pager.find(x => x.page === page)?.url!;
    const html = await this.getHtml(url);
    return this.parseMangaList(html);
  }

  parseMangaList(html: string): MangaList {
    let result = html.split("<div class=\"cat2_list text-center mb-4\">")
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

    let pager: Pager[] = [];
    if (html.includes('_pager">')) {
      result = html.split('_pager">')[1].split('</div>')[0].split('document.location.href=\'');
      result.shift();
      pager = result.map(x => {
        const url = x.split('\'">')[0];
        let page = x.split('\'">')[1].split('</li>')[0];
        if (html.includes('search_pager')) {
          if (page.includes('pageractive')) {
            page = page.split('pageractive">')[1].split('</a>')[0];
          }
        } else {
          page = page.replace('<font color="">', '').replace('</font>', '')
        }

        return {
          page,
          url,
        }
      })
    }

    return {manga, pager};
  }
}

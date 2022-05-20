import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { firstValueFrom } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MangaPictureParserService {

  constructor(
    private http: HttpClient
  ) { }

  getFirstPageUrl(comicId: string, chNo: string): any {
    // return `https://comicbus.live/online/a-${comicId}.html?ch=${chNo}`; // first version
    // return `https://comic.aya.click/online/b-${comicId}.html?ch=${chNo}`; // v20210405
    // return `https://comic.aya.click/online/best_${comicId}.html?ch=${chNo}`; // v20210627
    return `http://localhost:4200/online/new-${comicId}.html?ch=${chNo}`; // v20211113
  }

  async getHtml(comicId: string, chNo: string): Promise<string> {
    const url = this.getFirstPageUrl(comicId, chNo);
    return firstValueFrom(this.http.get(url, {responseType: "text"}));
  }

  async getHtmlPic(comicId: string, chNo: string): Promise<any> {
    const results = await this.getHtml(comicId, chNo);
    const title = results.split("<title>")[1].split("</title>")[0].split(" ")[0];
    const itemId = results.split("var ti=")[1].split(";")[0];
    const code = results.split("var n06yw_='")[1].split("'")[0];
    const url = this.getComicUrls(code, itemId)
    return url;
  }

  private getComicUrls(code: string, itemId: string) {
    const result = [];
    const partLength = 48;

    for (let keyIndex = 0; keyIndex < code.length; keyIndex += partLength) {
      // this.mySubstring(code, keyIndex + 6, null);
      const vol = this.decode(code.substring(keyIndex, keyIndex + 2));
      const urlPrefix = this.decode(code.substring(keyIndex + 2, keyIndex + 4));
      const pages = this.decode(code.substring(keyIndex + 4, keyIndex + 6));
      const subKey = code.substring(keyIndex + 6, keyIndex + 46);

      // const vol = this.getOnlyDigit(subKey.substr(0, 4));
      // const pages = parseInt(this.getOnlyDigit(subKey.substr(7, 3)), 0);

      const comicUrls = [];
      for (let page = 1; page <= pages; ++page) {
        comicUrls.push(this.urlCreator(subKey, urlPrefix, itemId, vol, page));
      }
      result.push({ Vol: this.padDigits(Number(vol), 4), Urls: comicUrls });
    }
    return result;
  }

  private urlCreator(subKey: string, urlPrefix: string, itemId: string, vol: string, page: number) {
    let sid = String(urlPrefix).substring(0,1);
    let did = String(urlPrefix).substring(1,2);
    const hash = subKey.substring(this.getHash(page), this.getHash(page) + 3)

    return 'http://img' + sid + '.8comic.com/' + did + '/' + itemId + '/' + vol + '/' + this.padDigits(page, 3) + '_' + hash + '.jpg';
  }

  private getHash(n: number) {
    return (((n - 1) / 10) % 10) + (((n - 1) % 10) * 3);
  }

  private padDigits(number: number, digits: number) {
    return Array(Math.max(digits - String(number).length + 1, 0)).join('0') + number;
  }

  private decode(code:string): any {
    if (code.length != 2)
      return code;
    var az = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
    var a = code.substring(0, 1);
    var b = code.substring(1, 2);
    if (a == "Z")
      return 8000 + az.indexOf(b);
    else
      return az.indexOf(a) * 52 + az.indexOf(b);
  }
}

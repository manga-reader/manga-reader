import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { firstValueFrom } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MangaImageParserService {

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

  async getHtmlImg(comicId: string, chNo: string): Promise<any> {
    const results = await this.getHtml(comicId, chNo);
    const title = results.split("<title>")[1].split("</title>")[0].split(" ")[0];
    const itemId = results.split("var ti=")[1].split(";")[0];   // comic id on url
    const codeVariableName = results.split("=lc(")[1].split(");")[0];   // string including variable name of code
    const codeVarName = codeVariableName.split("(")[1].split(",")[0];
    const code = results.split("var " + codeVarName + "='")[1].split("'")[0];   // hash code of manga images
    const partLength = this.getPartLength(results);

    const indexString = results.split(";i++){")[2].split("};")[0];
    const volVarName = indexString.split("== ch)")[0].split("if(")[1];
    const volIndex = this.parseIndex(indexString, volVarName);
    const urlPrefixVarName = String(indexString.split(", 0, 1)")[0].split("(").pop());
    const urlPrefixIndex = this.parseIndex(indexString, urlPrefixVarName);
    const pagesVarName = indexString.split(";ps=")[1].split(";")[0];
    const pagesIndex = this.parseIndex(indexString, pagesVarName);
    const subkeyVarName = String(indexString.split(",mm(p),3)")[0].split("(").pop());
    const subkeyIndex = this.parseIndex(indexString, subkeyVarName);
    const index = {
      vol: volIndex,
      urlPrefix: urlPrefixIndex,
      pages: pagesIndex,
      subKey: subkeyIndex,
    }

    const url = this.getComicUrls(code, itemId, partLength, index)
    return url;
  }

  private getComicUrls(code: string, itemId: string, partLength: number, index: any) {
    const result = [];

    for (let keyIndex = 0; keyIndex < code.length; keyIndex += partLength) {
      const vol = this.decode(code.substring(keyIndex + index.vol, keyIndex + index.vol + 2));
      const urlPrefix = this.decode(code.substring(keyIndex + index.urlPrefix, keyIndex + index.urlPrefix + 2));
      const pages = this.decode(code.substring(keyIndex + index.pages, keyIndex + index.pages + 2));
      const subKey = code.substring(keyIndex + index.subKey, keyIndex + index.subKey + 40);

      const comicUrls = [];
      for (let page = 1; page <= pages; ++page) {
        comicUrls.push(this.urlCreator(subKey, urlPrefix, itemId, vol, page));
      }
      result.push({ vol, Urls: comicUrls });
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

  private getPartLength(source:string): number {
    const partLengthVarName = source.split(",i*(")[1].split("+")[0].split("-")[0];
    const calculateStr = source.split(",i*(" + partLengthVarName)[1].split(")")[0];
    const sign = calculateStr.includes("+") ? "+" : "-";
    const value = calculateStr.split(sign)[1];
    return Number(source.split("var " + partLengthVarName + "=")[1].split(";")[0]) + Number(value) * (sign == "+" ? 1 : -1);
  }

  private parseIndex(source:string, varName:string): number {
    return Number(source.split("var " + varName + "=")[1].split("));")[0].split(",2")[0].split("+").pop());
  }
}

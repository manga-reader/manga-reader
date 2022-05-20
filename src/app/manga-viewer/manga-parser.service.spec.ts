import { TestBed } from '@angular/core/testing';

import { MangaPictureParserService } from './manga-picture-parser.service';

describe('MangaParserService', () => {
  let service: MangaPictureParserService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MangaPictureParserService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

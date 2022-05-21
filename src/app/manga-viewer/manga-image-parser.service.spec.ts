import { TestBed } from '@angular/core/testing';

import { MangaImageParserService } from './manga-image-parser.service';

describe('MangaImageParserService', () => {
  let service: MangaImageParserService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MangaImageParserService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

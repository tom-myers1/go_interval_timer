import { IntervalTimerPage } from './app.po';

describe('interval-timer App', function() {
  let page: IntervalTimerPage;

  beforeEach(() => {
    page = new IntervalTimerPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});

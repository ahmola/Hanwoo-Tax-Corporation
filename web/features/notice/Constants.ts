export interface NoticeItem {
  id: number;
  title: string;
  date: string;
  category: string;
  isImportant: boolean;
  content: string; // 본문 내용 추가
}

// API 개발 전 임시용
export const NOTICE_ITEMS: NoticeItem[] = [
  {
    id: 1,
    title: "2026년 1기 부가가치세 예정신고 안내",
    date: "2026.04.01",
    category: "세무일정",
    isImportant: true,
    content: `
      2026년 1기 부가가치세 예정신고 기간이 다가왔습니다.
      법인 사업자는 4월 25일까지 신고 및 납부를 완료해야 합니다.
      
      [필요 서류]
      1. 매출/매입 세금계산서 합계표
      2. 신용카드 매출전표 발행 금액 집계표
      ...
    `
  },
  {
    id: 2,
    title: "법인세법 개정에 따른 주요 변경사항 정리",
    date: "2026.03.15",
    category: "뉴스",
    isImportant: false,
    content: "올해 개정된 법인세법의 주요 내용은 다음과 같습니다..."
  },
  // ... 나머지 데이터도 content 추가
];
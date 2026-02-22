import NoticeDetailPage from "@/features/notice/NoticeDetailPage";

export const dynamic = 'force-dynamic';

export default async function Page({ params }: { params: Promise<{ id: string }> }) {
  // Next.js 라우터에서 넘겨주는 id값을 기능 컴포넌트에 전달만 함
  const {id} = await params;

  return <NoticeDetailPage id={id} />;
}
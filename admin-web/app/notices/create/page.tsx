import NoticeWriteForm from "@/features/notices/NoticeWriteForm";

export default function NoticeCreatePage() {
  return (
    <div className="table-primary">
      <h1 className="h1-primary">새 공지사항 작성</h1>
      <NoticeWriteForm />
    </div>
  );
}
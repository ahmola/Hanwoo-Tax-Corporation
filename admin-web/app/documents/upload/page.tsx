import DocumentUploadForm from "@/features/documents/DocumentUploadForm";

export default function DocumentUploadPage() {
  return (
    <div className="table-primary">
      <h1 className="h1-primary">자료실 파일 업로드</h1>
      <DocumentUploadForm />
    </div>
  );
}
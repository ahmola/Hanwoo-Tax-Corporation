import ContactsList from "@/features/contacts/ContactsListForm";

export default function ConsultationPage() {
  return (
    <div className="table-primary">
      <div className="flex justify-between items-center mb-6">
        <h1 className="h1-primary">상담 신청 관리</h1>
        <span className="bg-blue-100 text-blue-800 text-xs font-bold px-2 py-1 rounded">
          Live Data
        </span>
      </div>
      <ContactsList />
    </div>
  );
}
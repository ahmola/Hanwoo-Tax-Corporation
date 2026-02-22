// 메인 홈페이지
import ContactForm from "@/features/contact/ContactForm";
import DocumentsForm from "@/features/documents/DocumentsForm";
import FaqForm from "@/features/faq/FaqForm";
import HeroForm from "@/features/hero/HeroForm";
import LocationForm from "@/features/location/LocationForm";
import NoticeForm from "@/features/notice/NoticeForm";
import ServiceForm from "@/features/servicees/ServiceForm";

export const dynamic = 'force-dynamic';

export default function Home() {
  return (
    <div className="flex flex-col gap-20 pb-20">
      {/* Hero Section */}
      <HeroForm />

      {/* Services Section */}
      <ServiceForm />

      {/* Notice Section */}
      <NoticeForm />

      {/* Documents Section */}
      <DocumentsForm />

      {/* FAQ Section */}
      <FaqForm />

      {/* Location Section */}
      <LocationForm />

      {/* Contact Section */}
      <ContactForm />
    </div>
  );
}
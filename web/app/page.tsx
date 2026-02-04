// 메인 홈페이지
import ContactForm from "@/features/contact/ContactForm";
import HeroForm from "@/features/hero/HeroForm";
import ServiceForm from "@/features/servicees/ServiceForm";

export default function Home() {
  return (
    <div className="flex flex-col gap-20 pb-20">
      {/* Hero Section */}
      <HeroForm />

      {/* Services Section */}
      <ServiceForm />

      {/* Notice Section */}

      {/* Documents Section */}

      {/* FAQ Section */}

      {/* Location Section */}

      {/* Contact Section */}
      <ContactForm />
    </div>
  );
}
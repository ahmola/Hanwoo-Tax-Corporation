import { SERVICE_ITEMS } from "./Constants";

export default function ServiceForm() {
    return (
    <section id="services" className="container mx-auto px-4">
      <h2 className="text-3xl font-bold text-center mb-12 text-white-900">
        주요 서비스
      </h2>
      
      <div className="grid md:grid-cols-3 gap-8">
        {SERVICE_ITEMS.map((item, idx) => (
          <div 
            key={idx} 
            className="p-6 border rounded-xl hover:shadow-lg transition bg-white group hover:-translate-y-1 duration-300"
          >
            <item.icon className="w-10 h-10 text-blue-700 mb-4 group-hover:scale-110 transition" />
            <h3 className="text-xl font-bold mb-2 text-slate-800">
              {item.title}
            </h3>
            <p className="text-slate-600 leading-relaxed">
              {item.desc}
            </p>
          </div>
        ))}
      </div>
    </section>
  );
}
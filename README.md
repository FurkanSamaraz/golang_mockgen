# install
go get  github.com/golang/mock/mockgen/model
go install github.com/golang/mock/mockgen@v1.6.0

# Go Mockgen
Go dilinde, testlerde kullanılan mock nesneleri, gerçek nesnelerin davranışlarını taklit eden, ancak testleri daha tahmin edilebilir hale getirmek için kontrol edilebilir ve önceden belirlenebilir sonuçlar üreten nesnelerdir. Bu, birimi (unit) testlerinizi gerçek nesnelerle değil, mock nesneleriyle yürütmenizi sağlar, böylece testlerinizin dışa bağımlılıklardan etkilenme olasılığı azalır.

# Mock Nesne Oluşturmanın Genel Adımları:

# Arayüz Tasarımı: 
Öncelikle mock nesnenizi uygulayacağınız bir arayüz (interface) tanımlarsınız. Bu arayüz, gerçek nesnenin uygulamasını yansıtmalıdır.

# Mock Nesne Oluşturma: 
Bu arayüzü (interface) uygulayan bir mock nesne yapısı oluşturursunuz. Bu yapı, önceden belirlenmiş sonuçları döndürmek veya belirli davranışları simüle etmek için gerekli metotları içerir.

# Test Fonksiyonları: 
Gerçek testleri yazarken mock nesnelerinizi kullanırsınız. Bu mock nesneler, beklenen davranışları ve sonuçları sağlayarak testinizi oluşturur.


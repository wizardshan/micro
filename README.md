����https://prometheus.io/download/ ����prometheus-2.53.1.darwin-amd64.tar.gz ��ѹ����������

���У�./prometheus --config.file=prometheus.yml

http://localhost:9090



��װpushgateway

�ӹ���https://prometheus.io/download/ ����pushgateway-1.9.0.darwin-amd64.tar.gz ��ѹ����������
./pushgateway

http://localhost:9091

echo 'test_metric_1m{a="a", b="b"} 3.15' | curl --data-binary @- http://localhost:9091/metrics/job/testJob
echo 'test_metric_1m{a="a", b="c"} 3.16' | curl --data-binary @- http://localhost:9091/metrics/job/testJob


�ֲ�ʽ��·׷��
https://www.jaegertracing.io/download/
./jaeger-all-in-one  all-in-one ���ݴ����ڴ��У�����ѧϰ


��·׷�������ߣ� Tracer Provider ��

��Դ�� Resource ��

�������� Exporter ��


������� OpenTelemetry �е� Tracer Name �� Service Name �����໥��ϵ

����
�ڷֲ�ʽϵͳ�Ĺ۲��У�OpenTelemetry �ṩ��ǿ��Ĺ��ߺͿ�����ռ��ͷ���ң�����ݡ���� OpenTelemetry �еĹؼ������ Tracer Name �� Service Name��������Ч���������ܷǳ���Ҫ�����Ľ���ϸ���������������Լ�����֮��Ĺ�ϵ��


1. Tracer Name �ĸ���
   Tracer Name �� OpenTelemetry �����ڱ�ʶ����ң�����ݵ� Tracer ʵ������ͨ�����ض��Ŀ��Ӧ�ó���ģ���������Tracer Name ����Ҫ�����ǣ�

ģ���ʶ������ʶ��ң����������Ӧ�ó�����ĸ����ֻ�����ɵġ�
���ڹ���͵��ԣ��ڸ��ӵ�Ӧ�ó����У���ͬģ�������Ҫ������׷�ٲ��ԣ�Tracer Name �ڴ������˹ؼ����á�
2. Service Name �ĸ���
   Service Name �� OpenTelemetry ��������ʶ���������Ӧ�ó�������ơ���΢����ܹ��У�ÿ������һ�㶼���Լ��� Service Name��Service Name ����Ҫ��;������

����ʶ����΢����ܹ��У�Service Name �������ֲ�ͬ�ķ���
���ݾۺϣ��ڼ�غͷ��������У�Service Name ���ھۺϺͶԱ����Բ�ͬ�����ң�����ݡ�
3. Tracer Name �� Service Name �Ĺ�ϵ
   ��Ȼ Tracer Name �� Service Name �����ڱ�ʶ������ң�����ݵ���Դ����������Ӧ���а��ݲ�ͬ�Ľ�ɫ��

�㼶���죺Service Name ͨ����ʾ���߲㼶��Ӧ�ó������񣬶� Tracer Name ����������Ӧ�ó����ڵľ���ģ���⡣
�����ԣ�Tracer Name ͨ������ϸ���ȵ�׷�٣�����ģ�鼶�𣬶� Service Name ���ڱ�ʾ���������Ӧ�ó�������ܺ�״̬��
4. ʵ��Ӧ���е�����
   ��һ������ƽ̨��΢����ܹ��У�Service Name ������ payment-service��������֧���ķ��񡣶�����������У�������һ���������ÿ������ģ�飬�� Tracer Name ������ credit-card-processing������������У�Service Name ���ڱ�ʶ����֧�����񣬶� Tracer Name ���ڱ�ʶ�����о����ģ�顣

����
�� OpenTelemetry �У�Tracer Name �� Service Name �����Ͳ���ң�����ݵĹؼ������ȷ������ʹ��������������ڹ�����Ч����ά�������ڼ�صķֲ�ʽϵͳ������Ҫ�������������������������ϵͳ�Ŀɹ۲��ԣ��Ӷ��Ż����ܺͿ��ٶ�λ���⡣
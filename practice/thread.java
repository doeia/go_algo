public class TestThread implements Runnable {
    @Override
    public void run() {
        for (int i = 0; i < 200; i++) {
            System.out.println(i);
        }
    }

    public static void main(String[] args) {
        TestThread testThread = new TestThread();

        new Thread(testThread).start();

        for (int i = 0; i < 1000; i++) {
            System.out.println(i);
        }
    }
}
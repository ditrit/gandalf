package com.ditrit.gandalf.bus.kafkajava;

import com.ditrit.gandalf.bus.kafkajava.sample.KafkaJavaCLR;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@EnableAutoConfiguration
public class KafkaJavaApplication implements CommandLineRunner {

	@Autowired
	private KafkaJavaCLR kafkaJavaCLR;

	public static void main(String[] args) {
		SpringApplication.run(KafkaJavaApplication.class, args);
	}

	@Override
	public void run(String... args) throws Exception {
		kafkaJavaCLR.sample();
	}
}

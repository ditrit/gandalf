package com.ditrit.gandalf.workflowengine.zeebejava;

import com.ditrit.gandalf.workflowengine.zeebejava.sample.ZeebeJavaCLR;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@EnableAutoConfiguration
public class ZeebeJavaApplication implements CommandLineRunner {

	@Autowired
	private ZeebeJavaCLR zeebeJavaCLR;

	public static void main(String[] args) {
		SpringApplication.run(ZeebeJavaApplication.class, args);
	}

	@Override
	public void run(String... args) throws Exception {
		zeebeJavaCLR.sample();
	}
}
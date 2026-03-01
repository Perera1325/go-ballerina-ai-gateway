import ballerina/http;
import ballerina/log;

service /ai on new http:Listener(9090) {

    resource function get process() returns json {

        log:printInfo("AI Orchestration Layer Triggered");

        int tokenUsage = 100; // simulate AI token usage
        decimal cost = tokenUsage * 0.002;

        return {
            message: "AI Response Generated",
            tokensUsed: tokenUsage,
            estimatedCostUSD: cost
        };
    }
}
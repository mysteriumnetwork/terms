package network.mysterium.terms;

import java.io.IOException;
import java.io.InputStream;

public class Terms {

    public static final String TERMS_BOUNTY_PILOT = "TERMS_BOUNTY_PILOT.md";
    public static final String TERMS_END_USER = "TERMS_END_USER.md";
    public static final String TERMS_EXIT_NODE = "TERMS_EXIT_NODE.md";
    public static final String WARRANTY = "WARRANTY.md";
    public static final String BUILD_PROPERTIES = "build.properties";

    public static String bountyPilotMD() {
        try {
            byte[] bytes = asBytes(TERMS_BOUNTY_PILOT);
            return new String(bytes);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    public static String endUserMD() {
        try {
            byte[] bytes = asBytes(TERMS_END_USER);
            return new String(bytes);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    public static String exitNodeMD() {
        try {
            byte[] bytes = asBytes(TERMS_EXIT_NODE);
            return new String(bytes);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    public static String warrantyMD() {
        try {
            byte[] bytes = asBytes(WARRANTY);
            return new String(bytes);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    public static String version() {
        try {
            byte[] bytes = asBytes(BUILD_PROPERTIES);
            return new String(bytes);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }


    private static byte[] asBytes(String fileName) throws IOException {
        InputStream stream = Terms.class.getClassLoader().getResourceAsStream(fileName);
        byte[] bytes = new byte[stream.available()];
        stream.read(bytes);
        return bytes;
    }
}

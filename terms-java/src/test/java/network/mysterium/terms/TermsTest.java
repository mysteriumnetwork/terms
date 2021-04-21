package network.mysterium.terms;


import org.junit.jupiter.api.Test;

import java.net.URI;
import java.nio.file.Files;
import java.nio.file.Paths;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

class TermsTest {

    @Test
    public void testTerms() throws Exception {
        assertEquals(read(Terms.TERMS_BOUNTY_PILOT), Terms.bountyPilotMD());
        assertEquals(read(Terms.TERMS_END_USER), Terms.endUserMD());
        assertEquals(read(Terms.TERMS_EXIT_NODE), Terms.exitNodeMD());
        assertEquals(read(Terms.WARRANTY), Terms.warrantyMD());
        assertNotNull(Terms.version());
    }

    private String read(String fileName) throws Exception {
        URI uri = this.getClass().getResource("/" + fileName).toURI();
        return new String(Files.readAllBytes(Paths.get(uri)));
    }
}
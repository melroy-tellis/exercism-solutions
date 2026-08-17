class Fighter {

    boolean isVulnerable() {
        return true;
    }

    int getDamagePoints(Fighter fighter) {
        return 1;
    }
}

// TODO: define the Warrior class
class Warrior extends Fighter {
    boolean isVulnerable() {
        return false;
    }

    int getDamagePoints(Fighter fighter) {
        if (fighter.isVulnerable()) {
            return 10;
        }
        return 6;
    }
    
    public String toString() {
        return "Fighter is a Warrior";
    }
}

// TODO: define the Wizard class
class Wizard extends Fighter {
    private boolean isSpellPrepared = false;

    void prepareSpell() {
        this.isSpellPrepared = true;
    }
    boolean isVulnerable() {
        return !this.isSpellPrepared;
    }

    int getDamagePoints(Fighter fighter) {
        if (this.isSpellPrepared) {
            return 12;
        }
        return 3;
    }
    
    public String toString() {
        return "Fighter is a Wizard";
    }
}

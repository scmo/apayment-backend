package db

import (
	"github.com/scmo/apayment-backend/models"
	"github.com/scmo/apayment-backend/services"
)

func Seed_Contributions() {
	//Only seed if table is empty
	if cnt, _ := services.CountContributions(); cnt > 0 {
		return
	}
	services.CreateContribution(&models.Contribution{Code: 5416, Name: "Beitrag für besonders Tierfreundliche Stallhaltungssysteme"})
	services.CreateContribution(&models.Contribution{Code: 5417, Name: "Beitrag für regelmässigen Auslauf im Freien"})
}

func Seed_ControlPoints() {

	//Only seed if table is empty
	if cnt, _ := services.CountControlCategories(); cnt > 0 {
		return
	}
	if cnt, _ := services.CountPointGroups(); cnt > 0 {
		return
	}
	if cnt, _ := services.CountControlPoints(); cnt > 0 {
		return
	}
	Seed_BTS_Rindergattung_Wasserbueffel()
	Seed_RAUS_Weidetiere()
}

func Seed_BTS_Rindergattung_Wasserbueffel() {
	contribution, _ := services.GetContributionByCode(5416)
	cc := models.ControlCategory{ControlCategoryId: "12.01_2017", Abbreviation: "BTS - Rindergattung und Wasserbüffel", Contribution: contribution}
	services.CreateControlCategory(&cc)

	pg1 := models.PointGroup{PointGroupId: "A1", Abbreviation: "Rinder - Milchkühe", PointGroup: "Rinder - Milchkühe", PointGroupCode: 1110, ControlCategory: &cc}
	services.CreatePointGroup(&pg1)

	cp1 := models.ControlPoint{ControlPointId: "01", Abbreviation: "Alle Tiere frei in Gruppen gehalten",
		ControlPoint: "Alle Tiere der Kategorie in Gruppen gehalten oder zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4", PointGroup: &pg1}
	services.CreateControlPoint(&cp1)
	cp2 := models.ControlPoint{ControlPointId: "02", Abbreviation: "Mind. 15 Lux Tageslicht im Stall",
		ControlPoint: "Alle Ställe, in denen sich die Tiere überwiegend aufhalten, verfügen über Tageslicht von mindestens 15 Lux Stärke (Kunstlicht zur Beurteilung ausschalten!). In Ruhe- und Rückugsbereichen ist eine geringere Beleuchtung zulässig.", PointGroup: &pg1}
	services.CreateControlPoint(&cp2)

	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von der Gruppenhaltung für weniger als 10% der Tiere", Points: 60, ControlPoint: &cp1})
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von der Gruppenhaltung für 10 oder mehr % der Tiere", Points: 110, ControlPoint: &cp1})
	services.CreateLack(&models.Lack{Name: "Etwas zu wenig Tageslicht", Points: 10, ControlPoint: &cp2})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig Tageslicht", Points: 110, ControlPoint: &cp2})

	// A2	Rinder - andere Kühe
	pg2 := models.PointGroup{PointGroupId: "A2", Abbreviation: "Rinder - andere Kühe", PointGroup: "Rinder - andere Kühe", PointGroupCode: 1150, ControlCategory: &cc}
	services.CreatePointGroup(&pg2)

	cpA2_1 := models.ControlPoint{ControlPointId: "01", Abbreviation: "Alle Tiere frei in Gruppen gehalten",
		ControlPoint: "Alle Tiere der Kategorie in Gruppen gehalten oder zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_1)
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von der Gruppenhaltung für weniger als 10% der Tiere", Points: 60, ControlPoint: &cpA2_1})
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von der Gruppenhaltung für 10 oder mehr % der Tiere", Points: 110, ControlPoint: &cpA2_1})

	cpA2_2 := models.ControlPoint{ControlPointId: "02", Abbreviation: "Mind. 15 Lux Tageslicht im Stall", ControlPoint: "Alle Ställe, in denen sich die Tiere überwiegend aufhalten, verfügen über Tageslicht von mindestens 15 Lux Stärke (Kunstlicht zur Beurteilung ausschalten!). In Ruhe- und Rückugsbereichen ist eine geringere Beleuchtung zulässig.", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_2)
	services.CreateLack(&models.Lack{Name: "Etwas zu wenig Tageslicht", Points: 10, ControlPoint: &cpA2_2})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig Tageslicht", Points: 110, ControlPoint: &cpA2_2})

	cpA2_3 := models.ControlPoint{ControlPointId: "03", Abbreviation: "Fress- und Tränkebereich: befestigter Boden", ControlPoint: "Befestigter Boden, mit oder ohne Perforierung\n Ausnahme: Abkalbebox und Krankenabteil", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_3)
	services.CreateLack(&models.Lack{Name: "Fress- und Tränkebereich: unbefestigter Boden", Points: 110, ControlPoint: &cpA2_3})

	cpA2_4 := models.ControlPoint{ControlPointId: "04", Abbreviation: "Alle Tiere haben dauernd (jeden Tag /24h) Zugang zu BTS-Liegebereich und nicht eingestreutem Bereich", ControlPoint: "Alle Tiere der Katgorie haben dauernd (jeden Tag* / während 24h**) Zugang zu einem BTS-konformen Liegebereich und einem nicht eingestreuten Bereich *) Alternative zwischen 1.4. und 30.11.: 24 h am Tag auf Weide **) zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4	", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_4)
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für weniger als 10% der Tiere", Points: 60, ControlPoint: &cpA2_4})
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für 10 oder mehr % der Tiere", Points: 110, ControlPoint: &cpA2_4})

	cpA2_5 := models.ControlPoint{ControlPointId: "05.1", Abbreviation: "Liegebereich in Boxen-Laufställen mit Liegematten", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_5)

	cpA2_6 := models.ControlPoint{ControlPointId: "05.1.1", Abbreviation: "Liegemattenfabrikat BTS-konform", ControlPoint: "Bewirtschafter kann BTS-Konformität nachweisen:\n - Beleg der Mattenlieferfirma mit Name, BVET-Bewilligungsnummer und Datum der Installation\n Falls Mattenfabrikat ohne öffentlich zugänglichen Prüfbericht: betriebsspezifischer Prüfbericht nach Anhang 6 Bst. C Ziff. 1.3.", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_6)
	services.CreateLack(&models.Lack{Name: "Liegemattenfabrikat nicht BTS-konform bei weniger als 10% der Boxen", Points: 60, ControlPoint: &cpA2_6})
	services.CreateLack(&models.Lack{Name: "Liegemattenfabrikat nicht BTS-konform bei 10 oder mehr % der Boxen", Points: 110, ControlPoint: &cpA2_6})

	cpA2_7 := models.ControlPoint{ControlPointId: "05.1.2", Abbreviation: "Alle Liegematten ausschliesslich mit gehäckseltem Stroh eingestreut", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_7)
	services.CreateLack(&models.Lack{Name: "Zu wenig BTS-konforme Einstreu", Points: 10, ControlPoint: &cpA2_7})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig BTS-konforme Einstreu", Points: 60, ControlPoint: &cpA2_7})
	services.CreateLack(&models.Lack{Name: "Keine BTS-konforme Einstreu", Points: 110, ControlPoint: &cpA2_7})


	cpA2_8 := models.ControlPoint{ControlPointId: "05.2", Abbreviation: "Liegebereich in allen anderen Laufställen	", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_8)

	cpA2_9 := models.ControlPoint{ControlPointId: "05.2.1", Abbreviation: "Liegebereich: Strohmatratze oder gleichwertiger Liegebereich", ControlPoint: "Liegebereich: Strohmatratze oder für das Tier gleichwertige Unterterlage (z.B. Sägemehlbett) / ohne Perforierung", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_9)
	services.CreateLack(&models.Lack{Name: "Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf weniger als 10% der Fläche", Points: 60, ControlPoint: &cpA2_9})
	services.CreateLack(&models.Lack{Name: "Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf 10 oder mehr % der Fläche", Points: 110, ControlPoint: &cpA2_9})

	/*
		A3 - Rinder - weibliche Tiere, über 365 Tage alt, bis zur ersten Abkalbung
	*/
	pg3 := models.PointGroup{PointGroupId: "A3",
		Abbreviation: "Rinder - weibliche Tiere, über 365 Tage alt, bis zur ersten Abkalbung",
		PointGroup:   "weibliche Tiere über 365 - 730 Tage alt, ohne Abkalbung", PointGroupCode: 1128, ControlCategory: &cc}
	services.CreatePointGroup(&pg3)

	cpA3_1 := models.ControlPoint{ControlPointId: "01", Abbreviation: "Alle Tiere frei in Gruppen gehalten", ControlPoint: "Alle Tiere der Kategorie in Gruppen gehalten oder zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_1)
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von der Gruppenhaltung für weniger als 10% der Tiere", Points: 60, ControlPoint: &cpA3_1})
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von der Gruppenhaltung für 10 oder mehr % der Tiere", Points: 110, ControlPoint: &cpA3_1})

	cpA3_2 := models.ControlPoint{ControlPointId: "02", Abbreviation: "Mind. 15 Lux Tageslicht im Stall", ControlPoint: "Alle Ställe, in denen sich die Tiere überwiegend aufhalten, verfügen über Tageslicht von mindestens 15 Lux Stärke (Kunstlicht zur Beurteilung ausschalten!). In Ruhe- und Rückugsbereichen ist eine geringere Beleuchtung zulässig.", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_2)
	services.CreateLack(&models.Lack{Name: "Etwas zu wenig Tageslicht", Points: 10, ControlPoint: &cpA3_2})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig Tageslicht", Points: 110, ControlPoint: &cpA3_2})

	cpA3_3 := models.ControlPoint{ControlPointId: "03", Abbreviation: "Fress- und Tränkebereich: befestigter Boden", ControlPoint: "Befestigter Boden, mit oder ohne Perforierung\n Ausnahme: Abkalbebox und Krankenabteil", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_3)
	services.CreateLack(&models.Lack{Name: "Fress- und Tränkebereich: unbefestigter Boden", Points: 110, ControlPoint: &cpA3_3})

	cpA3_4 := models.ControlPoint{ControlPointId: "04", Abbreviation: "Alle Tiere haben dauernd (jeden Tag /24h) Zugang zu BTS-Liegebereich und nicht eingestreutem Bereich", ControlPoint: "Alle Tiere der Katgorie haben dauernd (jeden Tag* / während 24h**) Zugang zu einem BTS-konformen Liegebereich und einem nicht eingestreuten Bereich *) Alternative zwischen 1.4. und 30.11.: 24 h am Tag auf Weide **) zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4	", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_4)
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für weniger als 10% der Tiere", Points: 60, ControlPoint: &cpA3_4})
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für 10 oder mehr % der Tiere", Points: 110, ControlPoint: &cpA3_4})
	cpA3_5 := models.ControlPoint{ControlPointId: "05.1", Abbreviation: "Liegebereich in Boxen-Laufställen mit Liegematten", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_5)
	cpA3_6 := models.ControlPoint{ControlPointId: "05.1.1", Abbreviation: "Liegemattenfabrikat BTS-konform", ControlPoint: "Bewirtschafter kann BTS-Konformität nachweisen:\n - Beleg der Mattenlieferfirma mit Name, BVET-Bewilligungsnummer und Datum der Installation\n Falls Mattenfabrikat ohne öffentlich zugänglichen Prüfbericht: betriebsspezifischer Prüfbericht nach Anhang 6 Bst. C Ziff. 1.3.", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_6)
	services.CreateLack(&models.Lack{Name: "Liegemattenfabrikat nicht BTS-konform bei weniger als 10% der Boxen", Points: 60, ControlPoint: &cpA3_6})
	services.CreateLack(&models.Lack{Name: "Liegemattenfabrikat nicht BTS-konform bei 10 oder mehr % der Boxen", Points: 110, ControlPoint: &cpA3_6})
	cpA3_7 := models.ControlPoint{ControlPointId: "05.1.2", Abbreviation: "Alle Liegematten ausschliesslich mit gehäckseltem Stroh eingestreut", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_7)
	services.CreateLack(&models.Lack{Name: "Zu wenig BTS-konforme Einstreu", Points: 10, ControlPoint: &cpA3_7})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig BTS-konforme Einstreu", Points: 60, ControlPoint: &cpA3_7})
	services.CreateLack(&models.Lack{Name: "Keine BTS-konforme Einstreu", Points: 110, ControlPoint: &cpA3_7})
	cpA3_8 := models.ControlPoint{ControlPointId: "05.2", Abbreviation: "Liegebereich in allen anderen Laufställen	", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_8)
	cpA3_9 := models.ControlPoint{ControlPointId: "05.2.1", Abbreviation: "Liegebereich: Strohmatratze oder gleichwertiger Liegebereich", ControlPoint: "Liegebereich: Strohmatratze oder für das Tier gleichwertige Unterterlage (z.B. Sägemehlbett) / ohne Perforierung", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_9)
	services.CreateLack(&models.Lack{Name: "Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf weniger als 10% der Fläche", Points: 60, ControlPoint: &cpA3_9})
	services.CreateLack(&models.Lack{Name: "Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf 10 oder mehr % der Fläche", Points: 110, ControlPoint: &cpA3_9})


	/*
		A4 = Rinder - weibliche Tiere, über 160-365 Tage alt
	*/
	pg4 := models.PointGroup{PointGroupId: "A4", Abbreviation: "Rinder - weibliche Tiere, über 160-365 Tage alt", PointGroup: "weibliche Tiere über 160 - 365 Tage alt", PointGroupCode: 1141, ControlCategory: &cc}
	services.CreatePointGroup(&pg4)

	cpA4_1 := models.ControlPoint{ControlPointId: "01", Abbreviation: "Alle Tiere frei in Gruppen gehalten", ControlPoint: "Alle Tiere der Kategorie in Gruppen gehalten oder zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_1)
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von der Gruppenhaltung für weniger als 10% der Tiere", Points: 60, ControlPoint: &cpA4_1})
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von der Gruppenhaltung für 10 oder mehr % der Tiere", Points: 110, ControlPoint: &cpA4_1})
	cpA4_2 := models.ControlPoint{ControlPointId: "02", Abbreviation: "Mind. 15 Lux Tageslicht im Stall", ControlPoint: "Alle Ställe, in denen sich die Tiere überwiegend aufhalten, verfügen über Tageslicht von mindestens 15 Lux Stärke (Kunstlicht zur Beurteilung ausschalten!). In Ruhe- und Rückugsbereichen ist eine geringere Beleuchtung zulässig.", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_2)
	services.CreateLack(&models.Lack{Name: "Etwas zu wenig Tageslicht", Points: 10, ControlPoint: &cpA4_2})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig Tageslicht", Points: 110, ControlPoint: &cpA4_2})
	cpA4_3 := models.ControlPoint{ControlPointId: "03", Abbreviation: "Fress- und Tränkebereich: befestigter Boden", ControlPoint: "Befestigter Boden, mit oder ohne Perforierung\n Ausnahme: Abkalbebox und Krankenabteil", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_3)
	services.CreateLack(&models.Lack{Name: "Fress- und Tränkebereich: unbefestigter Boden", Points: 110, ControlPoint: &cpA4_3})
	cpA4_4 := models.ControlPoint{ControlPointId: "04", Abbreviation: "Alle Tiere haben dauernd (jeden Tag /24h) Zugang zu BTS-Liegebereich und nicht eingestreutem Bereich", ControlPoint: "Alle Tiere der Katgorie haben dauernd (jeden Tag* / während 24h**) Zugang zu einem BTS-konformen Liegebereich und einem nicht eingestreuten Bereich *) Alternative zwischen 1.4. und 30.11.: 24 h am Tag auf Weide **) zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4	", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_4)
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für weniger als 10% der Tiere", Points: 60, ControlPoint: &cpA4_4})
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für 10 oder mehr % der Tiere", Points: 110, ControlPoint: &cpA4_4})
	cpA4_5 := models.ControlPoint{ControlPointId: "05.1", Abbreviation: "Liegebereich in Boxen-Laufställen mit Liegematten", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_5)
	cpA4_6 := models.ControlPoint{ControlPointId: "05.1.1", Abbreviation: "Liegemattenfabrikat BTS-konform", ControlPoint: "Bewirtschafter kann BTS-Konformität nachweisen:\n - Beleg der Mattenlieferfirma mit Name, BVET-Bewilligungsnummer und Datum der Installation\n Falls Mattenfabrikat ohne öffentlich zugänglichen Prüfbericht: betriebsspezifischer Prüfbericht nach Anhang 6 Bst. C Ziff. 1.3.", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_6)
	services.CreateLack(&models.Lack{Name: "Liegemattenfabrikat nicht BTS-konform bei weniger als 10% der Boxen", Points: 60, ControlPoint: &cpA4_6})
	services.CreateLack(&models.Lack{Name: "Liegemattenfabrikat nicht BTS-konform bei 10 oder mehr % der Boxen", Points: 110, ControlPoint: &cpA4_6})
	cpA4_7 := models.ControlPoint{ControlPointId: "05.1.2", Abbreviation: "Alle Liegematten ausschliesslich mit gehäckseltem Stroh eingestreut", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_7)
	services.CreateLack(&models.Lack{Name: "Zu wenig BTS-konforme Einstreu", Points: 10, ControlPoint: &cpA4_7})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig BTS-konforme Einstreu", Points: 60, ControlPoint: &cpA4_7})
	services.CreateLack(&models.Lack{Name: "Keine BTS-konforme Einstreu", Points: 110, ControlPoint: &cpA4_7})
	cpA4_8 := models.ControlPoint{ControlPointId: "05.2", Abbreviation: "Liegebereich in allen anderen Laufställen	", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_8)
	cpA4_9 := models.ControlPoint{ControlPointId: "05.2.1", Abbreviation: "Liegebereich: Strohmatratze oder gleichwertiger Liegebereich", ControlPoint: "Liegebereich: Strohmatratze oder für das Tier gleichwertige Unterterlage (z.B. Sägemehlbett) / ohne Perforierung", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_9)
	services.CreateLack(&models.Lack{Name: "Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf weniger als 10% der Fläche", Points: 60, ControlPoint: &cpA4_9})
	services.CreateLack(&models.Lack{Name: "Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf 10 oder mehr % der Fläche", Points: 110, ControlPoint: &cpA4_9})

	/*
		A6 - Rinder - männliche Tiere, über 730 Tage alt
	*/
	pg6 := models.PointGroup{PointGroupId: "A6", Abbreviation: "Rinder - männliche Tiere, über 730 Tage alt", PointGroup: "männliche Tiere, über 730 Tage alt", PointGroupCode: 1124, ControlCategory: &cc}
	services.CreatePointGroup(&pg6)

	cpA6_1 := models.ControlPoint{ControlPointId: "01", Abbreviation: "Alle Tiere frei in Gruppen gehalten", ControlPoint: "Alle Tiere der Kategorie in Gruppen gehalten oder zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_1)
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von der Gruppenhaltung für weniger als 10% der Tiere", Points: 60, ControlPoint: &cpA6_1})
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von der Gruppenhaltung für 10 oder mehr % der Tiere", Points: 110, ControlPoint: &cpA6_1})
	cpA6_2 := models.ControlPoint{ControlPointId: "02", Abbreviation: "Mind. 15 Lux Tageslicht im Stall", ControlPoint: "Alle Ställe, in denen sich die Tiere überwiegend aufhalten, verfügen über Tageslicht von mindestens 15 Lux Stärke (Kunstlicht zur Beurteilung ausschalten!). In Ruhe- und Rückugsbereichen ist eine geringere Beleuchtung zulässig.", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_2)
	services.CreateLack(&models.Lack{Name: "Etwas zu wenig Tageslicht", Points: 10, ControlPoint: &cpA6_2})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig Tageslicht", Points: 110, ControlPoint: &cpA6_2})
	cpA6_3 := models.ControlPoint{ControlPointId: "03", Abbreviation: "Fress- und Tränkebereich: befestigter Boden", ControlPoint: "Befestigter Boden, mit oder ohne Perforierung\n Ausnahme: Abkalbebox und Krankenabteil", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_3)
	services.CreateLack(&models.Lack{Name: "Fress- und Tränkebereich: unbefestigter Boden", Points: 110, ControlPoint: &cpA6_3})
	cpA6_4 := models.ControlPoint{ControlPointId: "04", Abbreviation: "Alle Tiere haben dauernd (jeden Tag /24h) Zugang zu BTS-Liegebereich und nicht eingestreutem Bereich", ControlPoint: "Alle Tiere der Katgorie haben dauernd (jeden Tag* / während 24h**) Zugang zu einem BTS-konformen Liegebereich und einem nicht eingestreuten Bereich *) Alternative zwischen 1.4. und 30.11.: 24 h am Tag auf Weide **) zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4	", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_4)
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für weniger als 10% der Tiere", Points: 60, ControlPoint: &cpA6_4})
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für 10 oder mehr % der Tiere", Points: 110, ControlPoint: &cpA6_4})
	cpA6_5 := models.ControlPoint{ControlPointId: "05.1", Abbreviation: "Liegebereich in Boxen-Laufställen mit Liegematten", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_5)
	cpA6_6 := models.ControlPoint{ControlPointId: "05.1.1", Abbreviation: "Liegemattenfabrikat BTS-konform", ControlPoint: "Bewirtschafter kann BTS-Konformität nachweisen:\n - Beleg der Mattenlieferfirma mit Name, BVET-Bewilligungsnummer und Datum der Installation\n Falls Mattenfabrikat ohne öffentlich zugänglichen Prüfbericht: betriebsspezifischer Prüfbericht nach Anhang 6 Bst. C Ziff. 1.3.", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_6)
	services.CreateLack(&models.Lack{Name: "Liegemattenfabrikat nicht BTS-konform bei weniger als 10% der Boxen", Points: 60, ControlPoint: &cpA6_6})
	services.CreateLack(&models.Lack{Name: "Liegemattenfabrikat nicht BTS-konform bei 10 oder mehr % der Boxen", Points: 110, ControlPoint: &cpA6_6})
	cpA6_7 := models.ControlPoint{ControlPointId: "05.1.2", Abbreviation: "Alle Liegematten ausschliesslich mit gehäckseltem Stroh eingestreut", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_7)
	services.CreateLack(&models.Lack{Name: "Zu wenig BTS-konforme Einstreu", Points: 10, ControlPoint: &cpA6_7})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig BTS-konforme Einstreu", Points: 60, ControlPoint: &cpA6_7})
	services.CreateLack(&models.Lack{Name: "Keine BTS-konforme Einstreu", Points: 110, ControlPoint: &cpA6_7})
	cpA6_8 := models.ControlPoint{ControlPointId: "05.2", Abbreviation: "Liegebereich in allen anderen Laufställen	", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_8)
	cpA6_9 := models.ControlPoint{ControlPointId: "05.2.1", Abbreviation: "Liegebereich: Strohmatratze oder gleichwertiger Liegebereich", ControlPoint: "Liegebereich: Strohmatratze oder für das Tier gleichwertige Unterterlage (z.B. Sägemehlbett) / ohne Perforierung", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_9)
	services.CreateLack(&models.Lack{Name: "Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf weniger als 10% der Fläche", Points: 60, ControlPoint: &cpA6_9})
	services.CreateLack(&models.Lack{Name: "Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf 10 oder mehr % der Fläche", Points: 110, ControlPoint: &cpA6_9})

	/*
		A7 - Rinder - männliche Tiere, über 365- 730 Tage alt
	*/
	pg7 := models.PointGroup{PointGroupId: "A7", Abbreviation: "Rinder - männliche Tiere, über 365- 730 Tage alt", PointGroup: "männliche Tiere, über 365 bis 730 Tage alt", PointGroupCode: 1129, ControlCategory: &cc}
	services.CreatePointGroup(&pg7)

	cpA7_1 := models.ControlPoint{ControlPointId: "01", Abbreviation: "Alle Tiere frei in Gruppen gehalten", ControlPoint: "Alle Tiere der Kategorie in Gruppen gehalten oder zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_1)
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von der Gruppenhaltung für weniger als 10% der Tiere", Points: 60, ControlPoint: &cpA7_1})
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von der Gruppenhaltung für 10 oder mehr % der Tiere", Points: 110, ControlPoint: &cpA7_1})
	cpA7_2 := models.ControlPoint{ControlPointId: "02", Abbreviation: "Mind. 15 Lux Tageslicht im Stall", ControlPoint: "Alle Ställe, in denen sich die Tiere überwiegend aufhalten, verfügen über Tageslicht von mindestens 15 Lux Stärke (Kunstlicht zur Beurteilung ausschalten!). In Ruhe- und Rückugsbereichen ist eine geringere Beleuchtung zulässig.", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_2)
	services.CreateLack(&models.Lack{Name: "Etwas zu wenig Tageslicht", Points: 10, ControlPoint: &cpA7_2})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig Tageslicht", Points: 110, ControlPoint: &cpA7_2})
	cpA7_3 := models.ControlPoint{ControlPointId: "03", Abbreviation: "Fress- und Tränkebereich: befestigter Boden", ControlPoint: "Befestigter Boden, mit oder ohne Perforierung\n Ausnahme: Abkalbebox und Krankenabteil", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_3)
	services.CreateLack(&models.Lack{Name: "Fress- und Tränkebereich: unbefestigter Boden", Points: 110, ControlPoint: &cpA7_3})
	cpA7_4 := models.ControlPoint{ControlPointId: "04", Abbreviation: "Alle Tiere haben dauernd (jeden Tag /24h) Zugang zu BTS-Liegebereich und nicht eingestreutem Bereich", ControlPoint: "Alle Tiere der Katgorie haben dauernd (jeden Tag* / während 24h**) Zugang zu einem BTS-konformen Liegebereich und einem nicht eingestreuten Bereich *) Alternative zwischen 1.4. und 30.11.: 24 h am Tag auf Weide **) zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4	", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_4)
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für weniger als 10% der Tiere", Points: 60, ControlPoint: &cpA7_4})
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für 10 oder mehr % der Tiere", Points: 110, ControlPoint: &cpA7_4})
	cpA7_5 := models.ControlPoint{ControlPointId: "05.1", Abbreviation: "Liegebereich in Boxen-Laufställen mit Liegematten", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_5)
	cpA7_6 := models.ControlPoint{ControlPointId: "05.1.1", Abbreviation: "Liegemattenfabrikat BTS-konform", ControlPoint: "Bewirtschafter kann BTS-Konformität nachweisen:\n - Beleg der Mattenlieferfirma mit Name, BVET-Bewilligungsnummer und Datum der Installation\n Falls Mattenfabrikat ohne öffentlich zugänglichen Prüfbericht: betriebsspezifischer Prüfbericht nach Anhang 6 Bst. C Ziff. 1.3.", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_6)
	services.CreateLack(&models.Lack{Name: "Liegemattenfabrikat nicht BTS-konform bei weniger als 10% der Boxen", Points: 60, ControlPoint: &cpA7_6})
	services.CreateLack(&models.Lack{Name: "Liegemattenfabrikat nicht BTS-konform bei 10 oder mehr % der Boxen", Points: 110, ControlPoint: &cpA7_6})
	cpA7_7 := models.ControlPoint{ControlPointId: "05.1.2", Abbreviation: "Alle Liegematten ausschliesslich mit gehäckseltem Stroh eingestreut", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_7)
	services.CreateLack(&models.Lack{Name: "Zu wenig BTS-konforme Einstreu", Points: 10, ControlPoint: &cpA7_7})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig BTS-konforme Einstreu", Points: 60, ControlPoint: &cpA7_7})
	services.CreateLack(&models.Lack{Name: "Keine BTS-konforme Einstreu", Points: 110, ControlPoint: &cpA7_7})
	cpA7_8 := models.ControlPoint{ControlPointId: "05.2", Abbreviation: "Liegebereich in allen anderen Laufställen	", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_8)
	cpA7_9 := models.ControlPoint{ControlPointId: "05.2.1", Abbreviation: "Liegebereich: Strohmatratze oder gleichwertiger Liegebereich", ControlPoint: "Liegebereich: Strohmatratze oder für das Tier gleichwertige Unterterlage (z.B. Sägemehlbett) / ohne Perforierung", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_9)
	services.CreateLack(&models.Lack{Name: "Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf weniger als 10% der Fläche", Points: 60, ControlPoint: &cpA7_9})
	services.CreateLack(&models.Lack{Name: "Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf 10 oder mehr % der Fläche", Points: 110, ControlPoint: &cpA7_9})
	/*
		A8 - Rinder - männliche Tiere, über 160-365 Tage alt
	*/
	pg8 := models.PointGroup{PointGroupId: "A8", Abbreviation: "Rinder - männliche Tiere, über 160-365 Tage alt", PointGroup: "männliche Tiere, über 160 bis 365 Tage alt", PointGroupCode: 1143, ControlCategory: &cc}
	services.CreatePointGroup(&pg8)
	cpA8_1 := models.ControlPoint{ControlPointId: "01", Abbreviation: "Alle Tiere frei in Gruppen gehalten", ControlPoint: "Alle Tiere der Kategorie in Gruppen gehalten oder zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_1)
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von der Gruppenhaltung für weniger als 10% der Tiere", Points: 60, ControlPoint: &cpA8_1})
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von der Gruppenhaltung für 10 oder mehr % der Tiere", Points: 110, ControlPoint: &cpA8_1})
	cpA8_2 := models.ControlPoint{ControlPointId: "02", Abbreviation: "Mind. 15 Lux Tageslicht im Stall", ControlPoint: "Alle Ställe, in denen sich die Tiere überwiegend aufhalten, verfügen über Tageslicht von mindestens 15 Lux Stärke (Kunstlicht zur Beurteilung ausschalten!). In Ruhe- und Rückugsbereichen ist eine geringere Beleuchtung zulässig.", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_2)
	services.CreateLack(&models.Lack{Name: "Etwas zu wenig Tageslicht", Points: 10, ControlPoint: &cpA8_2})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig Tageslicht", Points: 110, ControlPoint: &cpA8_2})
	cpA8_3 := models.ControlPoint{ControlPointId: "03", Abbreviation: "Fress- und Tränkebereich: befestigter Boden", ControlPoint: "Befestigter Boden, mit oder ohne Perforierung\n Ausnahme: Abkalbebox und Krankenabteil", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_3)
	services.CreateLack(&models.Lack{Name: "Fress- und Tränkebereich: unbefestigter Boden", Points: 110, ControlPoint: &cpA8_3})
	cpA8_4 := models.ControlPoint{ControlPointId: "04", Abbreviation: "Alle Tiere haben dauernd (jeden Tag /24h) Zugang zu BTS-Liegebereich und nicht eingestreutem Bereich", ControlPoint: "Alle Tiere der Katgorie haben dauernd (jeden Tag* / während 24h**) Zugang zu einem BTS-konformen Liegebereich und einem nicht eingestreuten Bereich *) Alternative zwischen 1.4. und 30.11.: 24 h am Tag auf Weide **) zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4	", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_4)
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für weniger als 10% der Tiere", Points: 60, ControlPoint: &cpA8_4})
	services.CreateLack(&models.Lack{Name: "Nicht zulässige Abweichung von dauernd Zugang zu BTS-konfromeme Liegebereich und zu nicht eingestreutem Bereich für 10 oder mehr % der Tiere", Points: 110, ControlPoint: &cpA8_4})
	cpA8_5 := models.ControlPoint{ControlPointId: "05.1", Abbreviation: "Liegebereich in Boxen-Laufställen mit Liegematten", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_5)
	cpA8_6 := models.ControlPoint{ControlPointId: "05.1.1", Abbreviation: "Liegemattenfabrikat BTS-konform", ControlPoint: "Bewirtschafter kann BTS-Konformität nachweisen:\n - Beleg der Mattenlieferfirma mit Name, BVET-Bewilligungsnummer und Datum der Installation\n Falls Mattenfabrikat ohne öffentlich zugänglichen Prüfbericht: betriebsspezifischer Prüfbericht nach Anhang 6 Bst. C Ziff. 1.3.", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_6)
	services.CreateLack(&models.Lack{Name: "Liegemattenfabrikat nicht BTS-konform bei weniger als 10% der Boxen", Points: 60, ControlPoint: &cpA8_6})
	services.CreateLack(&models.Lack{Name: "Liegemattenfabrikat nicht BTS-konform bei 10 oder mehr % der Boxen", Points: 110, ControlPoint: &cpA8_6})
	cpA8_7 := models.ControlPoint{ControlPointId: "05.1.2", Abbreviation: "Alle Liegematten ausschliesslich mit gehäckseltem Stroh eingestreut", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_7)
	services.CreateLack(&models.Lack{Name: "Zu wenig BTS-konforme Einstreu", Points: 10, ControlPoint: &cpA8_7})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig BTS-konforme Einstreu", Points: 60, ControlPoint: &cpA8_7})
	services.CreateLack(&models.Lack{Name: "Keine BTS-konforme Einstreu", Points: 110, ControlPoint: &cpA8_7})
	cpA8_8 := models.ControlPoint{ControlPointId: "05.2", Abbreviation: "Liegebereich in allen anderen Laufställen	", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_8)
	cpA8_9 := models.ControlPoint{ControlPointId: "05.2.1", Abbreviation: "Liegebereich: Strohmatratze oder gleichwertiger Liegebereich", ControlPoint: "Liegebereich: Strohmatratze oder für das Tier gleichwertige Unterterlage (z.B. Sägemehlbett) / ohne Perforierung", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_9)
	services.CreateLack(&models.Lack{Name: "Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf weniger als 10% der Fläche", Points: 60, ControlPoint: &cpA8_9})
	services.CreateLack(&models.Lack{Name: "Liegebereich nicht BTS-konform (z.B. nicht kompakt) auf 10 oder mehr % der Fläche", Points: 110, ControlPoint: &cpA8_9})
}

func Seed_RAUS_Weidetiere() {
	contribution, _ := services.GetContributionByCode(5417)
	cc := models.ControlCategory{ControlCategoryId: "12.07_2017", Abbreviation: "RAUS-Weidetiere", Contribution: contribution}
	services.CreateControlCategory(&cc)

	// A1 -
	pg1 := models.PointGroup{PointGroupId: "A1", Abbreviation: "Milchkühe", PointGroupCode: 1110, ControlCategory: &cc}
	services.CreatePointGroup(&pg1)

	cpA1_1 := models.ControlPoint{ControlPointId: "01", Abbreviation: "Laufhof befindet sich im Freien", PointGroup: &pg1}
	services.CreateControlPoint(&cpA1_1)
	cpA1_2 := models.ControlPoint{ControlPointId: "02", Abbreviation: "Laufhöfe, die für das Rindvieh dauernd zugänglich sind: Gesamtfläche und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Stall-Laufhof-Skizze, die für die Berechnung der Gesamtfläche* und der ungedeckten Fläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte Gesamtfläche* und ungedeckte Fläche wurden korrekt berechnet (= nachrechnen) Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) * = den Tieren dauernd zugängliche Liege- + Fress- + Lauffläche inner- und ausserhalb des Stalls", PointGroup: &pg1}
	services.CreateControlPoint(&cpA1_2)
	cpA1_3 := models.ControlPoint{ControlPointId: "03", Abbreviation: "Alle übrigen Laufhöfe: Gesamte und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Laufhof-Skizze, die für die Berechnung der gesamten und der ungedeckten Laufhoffläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte gesamte und ungedeckte Laufhoffläche wurden korrekt berechnet (= nachrechnen) Kategorien A und B: Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) Kategorien C: mind. 25 % der Laufhoffläche müssen ungedeckt sein Kategorien D: mind. 50 % der Laufhoffläche müssen ungedeckt sein", PointGroup: &pg1}
	services.CreateControlPoint(&cpA1_3)
	cpA1_4 := models.ControlPoint{ControlPointId: "04", Abbreviation: "Die Laufhof-Skizze entspricht den Anforderungen, ist verifiziert und aktuell", ControlPoint: "Anforderungen: siehe 02 Aktuell: grober visueller Abgleich", PointGroup: &pg1}
	services.CreateControlPoint(&cpA1_4)
	cpA1_5 := models.ControlPoint{ControlPointId: "05", Abbreviation: "Aktuelle Tierzahl pro Auslauf kleiner oder gleich max. Tierzahl auf aktueller und verifizierter Laufhof-Skizze", ControlPoint: "Anzahl Tiere, denen aktuell gleichzeitig Auslauf gewährt wird: Tiere Zählen bzw. Selbstdeklaration des Landwirts verifizieren", PointGroup: &pg1}
	services.CreateControlPoint(&cpA1_5)
	cpA1_6 := models.ControlPoint{ControlPointId: "06", Abbreviation: "1.11.-28.2.: kein Schattennetz", PointGroup: &pg1}
	services.CreateControlPoint(&cpA1_6)
	cpA1_7 := models.ControlPoint{ControlPointId: "07", Abbreviation: "Unbefestigte Laufhöfe: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei unbefestigten Laufhöfen: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg1}
	services.CreateControlPoint(&cpA1_7)
	cpA1_8 := models.ControlPoint{ControlPointId: "08", Abbreviation: "Liegebereich ohne Perforierung und ausreichend mit geeigneter Einstreue bedeckt", PointGroup: &pg1}
	services.CreateControlPoint(&cpA1_8)
	cpA1_11 := models.ControlPoint{ControlPointId: "11", Abbreviation: "Weiden: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei Weiden: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg1}
	services.CreateControlPoint(&cpA1_11)
	cpA1_12 := models.ControlPoint{ControlPointId: "12", Abbreviation: "Weide kann an Weidetagen ca. 25 % des TS-Verzehrs decken", PointGroup: &pg1}
	services.CreateControlPoint(&cpA1_12)
	cpA1_141 := models.ControlPoint{ControlPointId: "14.1", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg1}
	services.CreateControlPoint(&cpA1_141)
	cpA1_1411 := models.ControlPoint{ControlPointId: "14.1.1", Abbreviation: "Auslauf-Dokumentation entspricht den Anforderungen", ControlPoint: "Anforderungen an die Dokumentation: 1. Auslauf je Auslaufgruppe bzw. je Tier eingetragen 2. a) Auslauf nach spätestens 3 Tagen eingetragen bzw. b) Anfang und Ende von Zeitspannen eingetragen, während denen die Tiere 24 h am Tag Zugang haben zu - einer Weide (01.05.-31.10.) bzw. - einem Laufhof (01.11.-30.04.) Ist die Einhaltung der Auslaufbestimmungen während des ganzen Jahres durch das Haltungssystem gewährleistet, ist kein Journal erforderlich", PointGroup: &pg1}
	services.CreateControlPoint(&cpA1_1411)
	cpA1_1412 := models.ControlPoint{ControlPointId: "14.1.2", Abbreviation: "genügend Weide- bzw. bis Auslauftage",
		ControlPoint: "Allen Tieren der Kategorie wurde vom 01.05 bis zum 31.10. an mindestens 26 Tagen pro Monat und vom 01.11. bis zum 30.04. an mindestens 13 Tagen pro Monat Auslauf auf einer Weide gewährt oder zulässige Abweichungen gemäss DZV", PointGroup: &pg1}
	services.CreateControlPoint(&cpA1_1412)
	cpA1_142 := models.ControlPoint{ControlPointId: "14.2", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg1}
	services.CreateControlPoint(&cpA1_142)
	cpA1_1421 := models.ControlPoint{ControlPointId: "14.2.1", Abbreviation: "Alternativ-Variante für betreffende Tiere zulässig und Laufhof dauernd zugänglich", ControlPoint: "Auslauf-Alternativ-Variante für Tiere der Rindergattung, - die gemästet werden; - männliche Zuchttiere; und - bis 160 Tage alte weibliche Zuchttiere Alle Tiere d. Kategorie hatten während des ganzen Jahres dauernd (24 h pro Tag) Zugang zu Laufhof oder zulässige Abweichungen gemäss DZV", PointGroup: &pg1}
	services.CreateControlPoint(&cpA1_1421)

	services.CreateLack(&models.Lack{Name: "Laufhof befindet sich nicht im Freien", Points: 110, ControlPoint: &cpA1_1})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA1_2})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA1_2})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA1_3})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA1_3})

	services.CreateLack(&models.Lack{Name: "Laufhof-Skizze kann nicht vorgewiesen werden, entspricht nicht den Anforderungen, ist nicht verifiziert oder nicht aktuell", Francs: 200, ControlPoint: &cpA1_4})

	services.CreateLack(&models.Lack{Name: "Aktuelle Tierzahl pro Auslauf grösser als maximale Tierzahl auf aktueller und verifizierter Laufhof-Skizze", Points: 110, ControlPoint: &cpA1_5})

	services.CreateLack(&models.Lack{Name: "Vom 1.11 ' 28.2. ein Schattennetz", Points: 10, ControlPoint: &cpA1_6})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA1_7})

	services.CreateLack(&models.Lack{Name: "Zu wenig geeignete Einstreu", Points: 10, ControlPoint: &cpA1_8})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig geeignete Einstreu", Points: 40, ControlPoint: &cpA1_8})
	services.CreateLack(&models.Lack{Name: "Keine geeignete Einstreu", Points: 110, ControlPoint: &cpA1_8})
	services.CreateLack(&models.Lack{Name: "Liegebereich(e) mit Perforation", Points: 110, ControlPoint: &cpA1_8})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA1_11})

	services.CreateLack(&models.Lack{Name: "Weide kann an Weidetagen nicht ca. 25% des TS-Verzehrs decken", Points: 60, ControlPoint: &cpA1_12})

	services.CreateLack(&models.Lack{Name: "Auslauf-Dokumentation entspricht nicht den Anforderungen", Francs: 200, ControlPoint: &cpA1_1411})

	services.CreateLack(&models.Lack{Name: "zu wenig Tage mit Zugang zur Weide- bzw. zum Laufhof nachgewiesen", Points: 4, Computed:true, ControlPoint: &cpA1_1412})

	services.CreateLack(&models.Lack{Name: "Alternativ Variante für betreffende Tiere nicht zulässig oder Laufhof nicht dauernd zugänglich", Points: 110, ControlPoint: &cpA1_1421})

	// A2	Rinder - andere Kühe
	pg2 := models.PointGroup{PointGroupId: "A2", Abbreviation: "Rinder - andere Kühe", PointGroup: "Rinder - andere Kühe", PointGroupCode: 1150, ControlCategory: &cc}
	services.CreatePointGroup(&pg2)

	cpA2_1 := models.ControlPoint{ControlPointId: "01", Abbreviation: "Laufhof befindet sich im Freien", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_1)
	cpA2_2 := models.ControlPoint{ControlPointId: "02", Abbreviation: "Laufhöfe, die für das Rindvieh dauernd zugänglich sind: Gesamtfläche und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Stall-Laufhof-Skizze, die für die Berechnung der Gesamtfläche* und der ungedeckten Fläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte Gesamtfläche* und ungedeckte Fläche wurden korrekt berechnet (= nachrechnen) Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) * = den Tieren dauernd zugängliche Liege- + Fress- + Lauffläche inner- und ausserhalb des Stalls", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_2)
	cpA2_3 := models.ControlPoint{ControlPointId: "03", Abbreviation: "Alle übrigen Laufhöfe: Gesamte und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Laufhof-Skizze, die für die Berechnung der gesamten und der ungedeckten Laufhoffläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte gesamte und ungedeckte Laufhoffläche wurden korrekt berechnet (= nachrechnen) Kategorien A und B: Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) Kategorien C: mind. 25 % der Laufhoffläche müssen ungedeckt sein Kategorien D: mind. 50 % der Laufhoffläche müssen ungedeckt sein", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_3)
	cpA2_4 := models.ControlPoint{ControlPointId: "04", Abbreviation: "Die Laufhof-Skizze entspricht den Anforderungen, ist verifiziert und aktuell", ControlPoint: "Anforderungen: siehe 02 Aktuell: grober visueller Abgleich", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_4)
	cpA2_5 := models.ControlPoint{ControlPointId: "05", Abbreviation: "Aktuelle Tierzahl pro Auslauf kleiner oder gleich max. Tierzahl auf aktueller und verifizierter Laufhof-Skizze", ControlPoint: "Anzahl Tiere, denen aktuell gleichzeitig Auslauf gewährt wird: Tiere Zählen bzw. Selbstdeklaration des Landwirts verifizieren", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_5)
	cpA2_6 := models.ControlPoint{ControlPointId: "06", Abbreviation: "1.11.-28.2.: kein Schattennetz", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_6)
	cpA2_7 := models.ControlPoint{ControlPointId: "07", Abbreviation: "Unbefestigte Laufhöfe: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei unbefestigten Laufhöfen: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_7)
	cpA2_8 := models.ControlPoint{ControlPointId: "08", Abbreviation: "Liegebereich ohne Perforierung und ausreichend mit geeigneter Einstreue bedeckt", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_8)
	cpA2_11 := models.ControlPoint{ControlPointId: "11", Abbreviation: "Weiden: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei Weiden: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_11)
	cpA2_12 := models.ControlPoint{ControlPointId: "12", Abbreviation: "Weide kann an Weidetagen ca. 25 % des TS-Verzehrs decken", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_12)
	cpA2_141 := models.ControlPoint{ControlPointId: "14.1", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_141)
	cpA2_1411 := models.ControlPoint{ControlPointId: "14.1.1", Abbreviation: "Auslauf-Dokumentation entspricht den Anforderungen", ControlPoint: "Anforderungen an die Dokumentation: 1. Auslauf je Auslaufgruppe bzw. je Tier eingetragen 2. a) Auslauf nach spätestens 3 Tagen eingetragen bzw. b) Anfang und Ende von Zeitspannen eingetragen, während denen die Tiere 24 h am Tag Zugang haben zu - einer Weide (01.05.-31.10.) bzw. - einem Laufhof (01.11.-30.04.) Ist die Einhaltung der Auslaufbestimmungen während des ganzen Jahres durch das Haltungssystem gewährleistet, ist kein Journal erforderlich", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_1411)
	cpA2_1412 := models.ControlPoint{ControlPointId: "14.1.2", Abbreviation: "genügend Weide- bzw. bis Auslauftage",
		ControlPoint: "Allen Tieren der Kategorie wurde vom 01.05 bis zum 31.10. an mindestens 26 Tagen pro Monat und vom 01.11. bis zum 30.04. an mindestens 13 Tagen pro Monat Auslauf auf einer Weide gewährt oder zulässige Abweichungen gemäss DZV", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_1412)
	cpA2_142 := models.ControlPoint{ControlPointId: "14.2", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_142)
	cpA2_1421 := models.ControlPoint{ControlPointId: "14.2.1", Abbreviation: "Alternativ-Variante für betreffende Tiere zulässig und Laufhof dauernd zugänglich", ControlPoint: "Auslauf-Alternativ-Variante für Tiere der Rindergattung, - die gemästet werden; - männliche Zuchttiere; und - bis 160 Tage alte weibliche Zuchttiere Alle Tiere d. Kategorie hatten während des ganzen Jahres dauernd (24 h pro Tag) Zugang zu Laufhof oder zulässige Abweichungen gemäss DZV", PointGroup: &pg2}
	services.CreateControlPoint(&cpA2_1421)

	services.CreateLack(&models.Lack{Name: "Laufhof befindet sich nicht im Freien", Points: 110, ControlPoint: &cpA2_1})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA2_2})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA2_2})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA2_3})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA2_3})

	services.CreateLack(&models.Lack{Name: "Laufhof-Skizze kann nicht vorgewiesen werden, entspricht nicht den Anforderungen, ist nicht verifiziert oder nicht aktuell", Francs: 200, ControlPoint: &cpA2_4})

	services.CreateLack(&models.Lack{Name: "Aktuelle Tierzahl pro Auslauf grösser als maximale Tierzahl auf aktueller und verifizierter Laufhof-Skizze", Points: 110, ControlPoint: &cpA2_5})

	services.CreateLack(&models.Lack{Name: "Vom 1.11 ' 28.2. ein Schattennetz", Points: 10, ControlPoint: &cpA2_6})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA2_7})

	services.CreateLack(&models.Lack{Name: "Zu wenig geeignete Einstreu", Points: 10, ControlPoint: &cpA2_8})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig geeignete Einstreu", Points: 40, ControlPoint: &cpA2_8})
	services.CreateLack(&models.Lack{Name: "Keine geeignete Einstreu", Points: 110, ControlPoint: &cpA2_8})
	services.CreateLack(&models.Lack{Name: "Liegebereich(e) mit Perforation", Points: 110, ControlPoint: &cpA2_8})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA2_11})
	services.CreateLack(&models.Lack{Name: "Weide kann an Weidetagen nicht ca. 25% des TS-Verzehrs decken", Points: 60, ControlPoint: &cpA2_12})

	services.CreateLack(&models.Lack{Name: "Auslauf-Dokumentation entspricht nicht den Anforderungen", Francs: 200, ControlPoint: &cpA2_1411})

	services.CreateLack(&models.Lack{Name: "zu wenig Tage mit Zugang zur Weide- bzw. zum Laufhof nachgewiesen", Points: 4, Computed:true, ControlPoint: &cpA2_1412})

	services.CreateLack(&models.Lack{Name: "Alternativ Variante für betreffende Tiere nicht zulässig oder Laufhof nicht dauernd zugänglich", Points: 110, ControlPoint: &cpA2_1421})

	// A3 - Rinder - weibliche Tiere, über 365 Tage alt, bis zur ersten Abkalbung
	pg3 := models.PointGroup{PointGroupId: "A3", Abbreviation: "Rinder - weibliche Tiere, über 365 Tage alt, bis zur ersten Abkalbung", PointGroup: "weibliche Tiere über 365 - 730 Tage alt, ohne Abkalbung", PointGroupCode: 1128, ControlCategory: &cc}
	services.CreatePointGroup(&pg3)

	cpA3_1 := models.ControlPoint{ControlPointId: "01", Abbreviation: "Laufhof befindet sich im Freien", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_1)
	cpA3_2 := models.ControlPoint{ControlPointId: "02", Abbreviation: "Laufhöfe, die für das Rindvieh dauernd zugänglich sind: Gesamtfläche und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Stall-Laufhof-Skizze, die für die Berechnung der Gesamtfläche* und der ungedeckten Fläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte Gesamtfläche* und ungedeckte Fläche wurden korrekt berechnet (= nachrechnen) Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) * = den Tieren dauernd zugängliche Liege- + Fress- + Lauffläche inner- und ausserhalb des Stalls", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_2)
	cpA3_3 := models.ControlPoint{ControlPointId: "03", Abbreviation: "Alle übrigen Laufhöfe: Gesamte und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Laufhof-Skizze, die für die Berechnung der gesamten und der ungedeckten Laufhoffläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte gesamte und ungedeckte Laufhoffläche wurden korrekt berechnet (= nachrechnen) Kategorien A und B: Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) Kategorien C: mind. 25 % der Laufhoffläche müssen ungedeckt sein Kategorien D: mind. 50 % der Laufhoffläche müssen ungedeckt sein", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_3)
	cpA3_4 := models.ControlPoint{ControlPointId: "04", Abbreviation: "Die Laufhof-Skizze entspricht den Anforderungen, ist verifiziert und aktuell", ControlPoint: "Anforderungen: siehe 02 Aktuell: grober visueller Abgleich", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_4)
	cpA3_5 := models.ControlPoint{ControlPointId: "05", Abbreviation: "Aktuelle Tierzahl pro Auslauf kleiner oder gleich max. Tierzahl auf aktueller und verifizierter Laufhof-Skizze", ControlPoint: "Anzahl Tiere, denen aktuell gleichzeitig Auslauf gewährt wird: Tiere Zählen bzw. Selbstdeklaration des Landwirts verifizieren", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_5)
	cpA3_6 := models.ControlPoint{ControlPointId: "06", Abbreviation: "1.11.-28.2.: kein Schattennetz", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_6)
	cpA3_7 := models.ControlPoint{ControlPointId: "07", Abbreviation: "Unbefestigte Laufhöfe: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei unbefestigten Laufhöfen: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_7)
	cpA3_8 := models.ControlPoint{ControlPointId: "08", Abbreviation: "Liegebereich ohne Perforierung und ausreichend mit geeigneter Einstreue bedeckt", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_8)
	cpA3_11 := models.ControlPoint{ControlPointId: "11", Abbreviation: "Weiden: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei Weiden: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_11)
	cpA3_12 := models.ControlPoint{ControlPointId: "12", Abbreviation: "Weide kann an Weidetagen ca. 25 % des TS-Verzehrs decken", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_12)
	cpA3_141 := models.ControlPoint{ControlPointId: "14.1", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_141)
	cpA3_1411 := models.ControlPoint{ControlPointId: "14.1.1", Abbreviation: "Auslauf-Dokumentation entspricht den Anforderungen", ControlPoint: "Anforderungen an die Dokumentation: 1. Auslauf je Auslaufgruppe bzw. je Tier eingetragen 2. a) Auslauf nach spätestens 3 Tagen eingetragen bzw. b) Anfang und Ende von Zeitspannen eingetragen, während denen die Tiere 24 h am Tag Zugang haben zu - einer Weide (01.05.-31.10.) bzw. - einem Laufhof (01.11.-30.04.) Ist die Einhaltung der Auslaufbestimmungen während des ganzen Jahres durch das Haltungssystem gewährleistet, ist kein Journal erforderlich", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_1411)
	cpA3_1412 := models.ControlPoint{ControlPointId: "14.1.2", Abbreviation: "genügend Weide- bzw. bis Auslauftage",
		ControlPoint: "Allen Tieren der Kategorie wurde vom 01.05 bis zum 31.10. an mindestens 26 Tagen pro Monat und vom 01.11. bis zum 30.04. an mindestens 13 Tagen pro Monat Auslauf auf einer Weide gewährt oder zulässige Abweichungen gemäss DZV", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_1412)
	cpA3_142 := models.ControlPoint{ControlPointId: "14.2", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_142)
	cpA3_1421 := models.ControlPoint{ControlPointId: "14.2.1", Abbreviation: "Alternativ-Variante für betreffende Tiere zulässig und Laufhof dauernd zugänglich", ControlPoint: "Auslauf-Alternativ-Variante für Tiere der Rindergattung, - die gemästet werden; - männliche Zuchttiere; und - bis 160 Tage alte weibliche Zuchttiere Alle Tiere d. Kategorie hatten während des ganzen Jahres dauernd (24 h pro Tag) Zugang zu Laufhof oder zulässige Abweichungen gemäss DZV", PointGroup: &pg3}
	services.CreateControlPoint(&cpA3_1421)

	services.CreateLack(&models.Lack{Name: "Laufhof befindet sich nicht im Freien", Points: 110, ControlPoint: &cpA3_1})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA3_2})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA3_2})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA3_3})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA3_3})

	services.CreateLack(&models.Lack{Name: "Laufhof-Skizze kann nicht vorgewiesen werden, entspricht nicht den Anforderungen, ist nicht verifiziert oder nicht aktuell", Francs: 200, ControlPoint: &cpA3_4})

	services.CreateLack(&models.Lack{Name: "Aktuelle Tierzahl pro Auslauf grösser als maximale Tierzahl auf aktueller und verifizierter Laufhof-Skizze", Points: 110, ControlPoint: &cpA3_5})

	services.CreateLack(&models.Lack{Name: "Vom 1.11 ' 28.2. ein Schattennetz", Points: 10, ControlPoint: &cpA3_6})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA3_7})

	services.CreateLack(&models.Lack{Name: "Zu wenig geeignete Einstreu", Points: 10, ControlPoint: &cpA3_8})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig geeignete Einstreu", Points: 40, ControlPoint: &cpA3_8})
	services.CreateLack(&models.Lack{Name: "Keine geeignete Einstreu", Points: 110, ControlPoint: &cpA3_8})
	services.CreateLack(&models.Lack{Name: "Liegebereich(e) mit Perforation", Points: 110, ControlPoint: &cpA3_8})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA3_11})

	services.CreateLack(&models.Lack{Name: "Weide kann an Weidetagen nicht ca. 25% des TS-Verzehrs decken", Points: 60, ControlPoint: &cpA3_12})

	services.CreateLack(&models.Lack{Name: "Auslauf-Dokumentation entspricht nicht den Anforderungen", Francs: 200, ControlPoint: &cpA3_1411})

	services.CreateLack(&models.Lack{Name: "zu wenig Tage mit Zugang zur Weide- bzw. zum Laufhof nachgewiesen", Points: 4, Computed:true, ControlPoint: &cpA3_1412})

	services.CreateLack(&models.Lack{Name: "Alternativ Variante für betreffende Tiere nicht zulässig oder Laufhof nicht dauernd zugänglich", Points: 110, ControlPoint: &cpA3_1421})

	// A4 = Rinder - weibliche Tiere, über 160-365 Tage alt
	pg4 := models.PointGroup{PointGroupId: "A4", Abbreviation: "Rinder - weibliche Tiere, über 160-365 Tage alt", PointGroup: "weibliche Tiere über 160 - 365 Tage alt", PointGroupCode: 1141, ControlCategory: &cc}
	services.CreatePointGroup(&pg4)

	cpA4_1 := models.ControlPoint{ControlPointId: "01", Abbreviation: "Laufhof befindet sich im Freien", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_1)
	cpA4_2 := models.ControlPoint{ControlPointId: "02", Abbreviation: "Laufhöfe, die für das Rindvieh dauernd zugänglich sind: Gesamtfläche und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Stall-Laufhof-Skizze, die für die Berechnung der Gesamtfläche* und der ungedeckten Fläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte Gesamtfläche* und ungedeckte Fläche wurden korrekt berechnet (= nachrechnen) Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) * = den Tieren dauernd zugängliche Liege- + Fress- + Lauffläche inner- und ausserhalb des Stalls", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_2)
	cpA4_3 := models.ControlPoint{ControlPointId: "03", Abbreviation: "Alle übrigen Laufhöfe: Gesamte und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Laufhof-Skizze, die für die Berechnung der gesamten und der ungedeckten Laufhoffläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte gesamte und ungedeckte Laufhoffläche wurden korrekt berechnet (= nachrechnen) Kategorien A und B: Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) Kategorien C: mind. 25 % der Laufhoffläche müssen ungedeckt sein Kategorien D: mind. 50 % der Laufhoffläche müssen ungedeckt sein", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_3)
	cpA4_4 := models.ControlPoint{ControlPointId: "04", Abbreviation: "Die Laufhof-Skizze entspricht den Anforderungen, ist verifiziert und aktuell", ControlPoint: "Anforderungen: siehe 02 Aktuell: grober visueller Abgleich", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_4)
	cpA4_5 := models.ControlPoint{ControlPointId: "05", Abbreviation: "Aktuelle Tierzahl pro Auslauf kleiner oder gleich max. Tierzahl auf aktueller und verifizierter Laufhof-Skizze", ControlPoint: "Anzahl Tiere, denen aktuell gleichzeitig Auslauf gewährt wird: Tiere Zählen bzw. Selbstdeklaration des Landwirts verifizieren", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_5)
	cpA4_6 := models.ControlPoint{ControlPointId: "06", Abbreviation: "1.11.-28.2.: kein Schattennetz", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_6)
	cpA4_7 := models.ControlPoint{ControlPointId: "07", Abbreviation: "Unbefestigte Laufhöfe: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei unbefestigten Laufhöfen: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_7)
	cpA4_8 := models.ControlPoint{ControlPointId: "08", Abbreviation: "Liegebereich ohne Perforierung und ausreichend mit geeigneter Einstreue bedeckt", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_8)
	cpA4_11 := models.ControlPoint{ControlPointId: "11", Abbreviation: "Weiden: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei Weiden: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_11)
	cpA4_12 := models.ControlPoint{ControlPointId: "12", Abbreviation: "Weide kann an Weidetagen ca. 25 % des TS-Verzehrs decken", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_12)
	cpA4_141 := models.ControlPoint{ControlPointId: "14.1", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_141)
	cpA4_1411 := models.ControlPoint{ControlPointId: "14.1.1", Abbreviation: "Auslauf-Dokumentation entspricht den Anforderungen", ControlPoint: "Anforderungen an die Dokumentation: 1. Auslauf je Auslaufgruppe bzw. je Tier eingetragen 2. a) Auslauf nach spätestens 3 Tagen eingetragen bzw. b) Anfang und Ende von Zeitspannen eingetragen, während denen die Tiere 24 h am Tag Zugang haben zu - einer Weide (01.05.-31.10.) bzw. - einem Laufhof (01.11.-30.04.) Ist die Einhaltung der Auslaufbestimmungen während des ganzen Jahres durch das Haltungssystem gewährleistet, ist kein Journal erforderlich", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_1411)
	cpA4_1412 := models.ControlPoint{ControlPointId: "14.1.2", Abbreviation: "genügend Weide- bzw. bis Auslauftage",
		ControlPoint: "Allen Tieren der Kategorie wurde vom 01.05 bis zum 31.10. an mindestens 26 Tagen pro Monat und vom 01.11. bis zum 30.04. an mindestens 13 Tagen pro Monat Auslauf auf einer Weide gewährt oder zulässige Abweichungen gemäss DZV", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_1412)
	cpA4_142 := models.ControlPoint{ControlPointId: "14.2", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_142)
	cpA4_1421 := models.ControlPoint{ControlPointId: "14.2.1", Abbreviation: "Alternativ-Variante für betreffende Tiere zulässig und Laufhof dauernd zugänglich", ControlPoint: "Auslauf-Alternativ-Variante für Tiere der Rindergattung, - die gemästet werden; - männliche Zuchttiere; und - bis 160 Tage alte weibliche Zuchttiere Alle Tiere d. Kategorie hatten während des ganzen Jahres dauernd (24 h pro Tag) Zugang zu Laufhof oder zulässige Abweichungen gemäss DZV", PointGroup: &pg4}
	services.CreateControlPoint(&cpA4_1421)

	services.CreateLack(&models.Lack{Name: "Laufhof befindet sich nicht im Freien", Points: 110, ControlPoint: &cpA4_1})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA4_2})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA4_2})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA4_3})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA4_3})

	services.CreateLack(&models.Lack{Name: "Laufhof-Skizze kann nicht vorgewiesen werden, entspricht nicht den Anforderungen, ist nicht verifiziert oder nicht aktuell", Francs: 200, ControlPoint: &cpA4_4})

	services.CreateLack(&models.Lack{Name: "Aktuelle Tierzahl pro Auslauf grösser als maximale Tierzahl auf aktueller und verifizierter Laufhof-Skizze", Points: 110, ControlPoint: &cpA4_5})

	services.CreateLack(&models.Lack{Name: "Vom 1.11 ' 28.2. ein Schattennetz", Points: 10, ControlPoint: &cpA4_6})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA4_7})

	services.CreateLack(&models.Lack{Name: "Zu wenig geeignete Einstreu", Points: 10, ControlPoint: &cpA4_8})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig geeignete Einstreu", Points: 40, ControlPoint: &cpA4_8})
	services.CreateLack(&models.Lack{Name: "Keine geeignete Einstreu", Points: 110, ControlPoint: &cpA4_8})
	services.CreateLack(&models.Lack{Name: "Liegebereich(e) mit Perforation", Points: 110, ControlPoint: &cpA4_8})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA4_11})

	services.CreateLack(&models.Lack{Name: "Weide kann an Weidetagen nicht ca. 25% des TS-Verzehrs decken", Points: 60, ControlPoint: &cpA4_12})

	services.CreateLack(&models.Lack{Name: "Auslauf-Dokumentation entspricht nicht den Anforderungen", Francs: 200, ControlPoint: &cpA4_1411})

	services.CreateLack(&models.Lack{Name: "zu wenig Tage mit Zugang zur Weide- bzw. zum Laufhof nachgewiesen", Points: 4, Computed:true, ControlPoint: &cpA4_1412})

	services.CreateLack(&models.Lack{Name: "Alternativ Variante für betreffende Tiere nicht zulässig oder Laufhof nicht dauernd zugänglich", Points: 110, ControlPoint: &cpA4_1421})

	// A5 - Rinder - weibliche Tiere, bis 160 Tage alt
	pg5 := models.PointGroup{PointGroupId: "A5", Abbreviation: "Rinder - weibliche Tiere, bis 160 Tage alt", PointGroup: "weibliche Tiere bis 160 Tage alt", PointGroupCode: 1142, ControlCategory: &cc}
	services.CreatePointGroup(&pg5)

	cpA5_1 := models.ControlPoint{ControlPointId: "01", Abbreviation: "Laufhof befindet sich im Freien", PointGroup: &pg5}
	services.CreateControlPoint(&cpA5_1)
	cpA5_2 := models.ControlPoint{ControlPointId: "02", Abbreviation: "Laufhöfe, die für das Rindvieh dauernd zugänglich sind: Gesamtfläche und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Stall-Laufhof-Skizze, die für die Berechnung der Gesamtfläche* und der ungedeckten Fläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte Gesamtfläche* und ungedeckte Fläche wurden korrekt berechnet (= nachrechnen) Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) * = den Tieren dauernd zugängliche Liege- + Fress- + Lauffläche inner- und ausserhalb des Stalls", PointGroup: &pg5}
	services.CreateControlPoint(&cpA5_2)
	cpA5_3 := models.ControlPoint{ControlPointId: "03", Abbreviation: "Alle übrigen Laufhöfe: Gesamte und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Laufhof-Skizze, die für die Berechnung der gesamten und der ungedeckten Laufhoffläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte gesamte und ungedeckte Laufhoffläche wurden korrekt berechnet (= nachrechnen) Kategorien A und B: Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) Kategorien C: mind. 25 % der Laufhoffläche müssen ungedeckt sein Kategorien D: mind. 50 % der Laufhoffläche müssen ungedeckt sein", PointGroup: &pg5}
	services.CreateControlPoint(&cpA5_3)
	cpA5_4 := models.ControlPoint{ControlPointId: "04", Abbreviation: "Die Laufhof-Skizze entspricht den Anforderungen, ist verifiziert und aktuell", ControlPoint: "Anforderungen: siehe 02 Aktuell: grober visueller Abgleich", PointGroup: &pg5}
	services.CreateControlPoint(&cpA5_4)
	cpA5_5 := models.ControlPoint{ControlPointId: "05", Abbreviation: "Aktuelle Tierzahl pro Auslauf kleiner oder gleich max. Tierzahl auf aktueller und verifizierter Laufhof-Skizze", ControlPoint: "Anzahl Tiere, denen aktuell gleichzeitig Auslauf gewährt wird: Tiere Zählen bzw. Selbstdeklaration des Landwirts verifizieren", PointGroup: &pg5}
	services.CreateControlPoint(&cpA5_5)
	cpA5_6 := models.ControlPoint{ControlPointId: "06", Abbreviation: "1.11.-28.2.: kein Schattennetz", PointGroup: &pg5}
	services.CreateControlPoint(&cpA5_6)
	cpA5_7 := models.ControlPoint{ControlPointId: "07", Abbreviation: "Unbefestigte Laufhöfe: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei unbefestigten Laufhöfen: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg5}
	services.CreateControlPoint(&cpA5_7)
	cpA5_8 := models.ControlPoint{ControlPointId: "08", Abbreviation: "Liegebereich ohne Perforierung und ausreichend mit geeigneter Einstreue bedeckt", PointGroup: &pg5}
	services.CreateControlPoint(&cpA5_8)
	cpA5_9 := models.ControlPoint{ControlPointId: "08", Abbreviation: "Bis 160 Tage alte Tiere nicht fixiert", PointGroup: &pg5}
	services.CreateControlPoint(&cpA5_9)
	cpA5_11 := models.ControlPoint{ControlPointId: "11", Abbreviation: "Weiden: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei Weiden: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg5}
	services.CreateControlPoint(&cpA5_11)
	cpA5_12 := models.ControlPoint{ControlPointId: "12", Abbreviation: "Weide kann an Weidetagen ca. 25 % des TS-Verzehrs decken", PointGroup: &pg5}
	services.CreateControlPoint(&cpA5_12)
	cpA5_141 := models.ControlPoint{ControlPointId: "14.1", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg5}
	services.CreateControlPoint(&cpA5_141)
	cpA5_1411 := models.ControlPoint{ControlPointId: "14.1.1", Abbreviation: "Auslauf-Dokumentation entspricht den Anforderungen", ControlPoint: "Anforderungen an die Dokumentation: 1. Auslauf je Auslaufgruppe bzw. je Tier eingetragen 2. a) Auslauf nach spätestens 3 Tagen eingetragen bzw. b) Anfang und Ende von Zeitspannen eingetragen, während denen die Tiere 24 h am Tag Zugang haben zu - einer Weide (01.05.-31.10.) bzw. - einem Laufhof (01.11.-30.04.) Ist die Einhaltung der Auslaufbestimmungen während des ganzen Jahres durch das Haltungssystem gewährleistet, ist kein Journal erforderlich", PointGroup: &pg5}
	services.CreateControlPoint(&cpA5_1411)
	cpA5_1412 := models.ControlPoint{ControlPointId: "14.1.2", Abbreviation: "genügend Weide- bzw. bis Auslauftage",
		ControlPoint: "Allen Tieren der Kategorie wurde vom 01.05 bis zum 31.10. an mindestens 26 Tagen pro Monat und vom 01.11. bis zum 30.04. an mindestens 13 Tagen pro Monat Auslauf auf einer Weide gewährt oder zulässige Abweichungen gemäss DZV", PointGroup: &pg5}
	services.CreateControlPoint(&cpA5_1412)
	cpA5_142 := models.ControlPoint{ControlPointId: "14.2", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg5}
	services.CreateControlPoint(&cpA5_142)
	cpA5_1421 := models.ControlPoint{ControlPointId: "14.2.1", Abbreviation: "Alternativ-Variante für betreffende Tiere zulässig und Laufhof dauernd zugänglich", ControlPoint: "Auslauf-Alternativ-Variante für Tiere der Rindergattung, - die gemästet werden; - männliche Zuchttiere; und - bis 160 Tage alte weibliche Zuchttiere Alle Tiere d. Kategorie hatten während des ganzen Jahres dauernd (24 h pro Tag) Zugang zu Laufhof oder zulässige Abweichungen gemäss DZV", PointGroup: &pg5}
	services.CreateControlPoint(&cpA5_1421)

	services.CreateLack(&models.Lack{Name: "Laufhof befindet sich nicht im Freien", Points: 110, ControlPoint: &cpA5_1})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA5_2})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA5_2})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA5_3})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA5_3})

	services.CreateLack(&models.Lack{Name: "Laufhof-Skizze kann nicht vorgewiesen werden, entspricht nicht den Anforderungen, ist nicht verifiziert oder nicht aktuell", Francs: 200, ControlPoint: &cpA5_4})

	services.CreateLack(&models.Lack{Name: "Aktuelle Tierzahl pro Auslauf grösser als maximale Tierzahl auf aktueller und verifizierter Laufhof-Skizze", Points: 110, ControlPoint: &cpA5_5})

	services.CreateLack(&models.Lack{Name: "Vom 1.11 ' 28.2. ein Schattennetz", Points: 10, ControlPoint: &cpA5_6})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA5_7})

	services.CreateLack(&models.Lack{Name: "Zu wenig geeignete Einstreu", Points: 10, ControlPoint: &cpA5_8})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig geeignete Einstreu", Points: 40, ControlPoint: &cpA5_8})
	services.CreateLack(&models.Lack{Name: "Keine geeignete Einstreu", Points: 110, ControlPoint: &cpA5_8})
	services.CreateLack(&models.Lack{Name: "Liegebereich(e) mit Perforation", Points: 110, ControlPoint: &cpA5_8})

	services.CreateLack(&models.Lack{Name: "Alter des fixierten Tiers weniger als 4 Monate", Points: 110, ControlPoint: &cpA5_9})
	services.CreateLack(&models.Lack{Name: "2015 und 2016: Alter des fixierten Tiers älter als 4 Monate", ControlPoint: &cpA5_9})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA5_11})

	services.CreateLack(&models.Lack{Name: "Weide kann an Weidetagen nicht ca. 25% des TS-Verzehrs decken", Points: 60, ControlPoint: &cpA5_12})

	services.CreateLack(&models.Lack{Name: "Auslauf-Dokumentation entspricht nicht den Anforderungen", Francs: 200, ControlPoint: &cpA5_1411})

	services.CreateLack(&models.Lack{Name: "zu wenig Tage mit Zugang zur Weide- bzw. zum Laufhof nachgewiesen", Points: 4, Computed:true, ControlPoint: &cpA5_1412})

	services.CreateLack(&models.Lack{Name: "Alternativ Variante für betreffende Tiere nicht zulässig oder Laufhof nicht dauernd zugänglich", Points: 110, ControlPoint: &cpA5_1421})

	// A6 - Rinder - männliche Tiere, über 730 Tage alt
	pg6 := models.PointGroup{PointGroupId: "A6", Abbreviation: "Rinder - männliche Tiere, über 730 Tage alt", PointGroup: "männliche Tiere, über 730 Tage alt", PointGroupCode: 1124, ControlCategory: &cc}
	services.CreatePointGroup(&pg6)

	cpA6_1 := models.ControlPoint{ControlPointId: "01", Abbreviation: "Laufhof befindet sich im Freien", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_1)
	cpA6_2 := models.ControlPoint{ControlPointId: "02", Abbreviation: "Laufhöfe, die für das Rindvieh dauernd zugänglich sind: Gesamtfläche und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Stall-Laufhof-Skizze, die für die Berechnung der Gesamtfläche* und der ungedeckten Fläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte Gesamtfläche* und ungedeckte Fläche wurden korrekt berechnet (= nachrechnen) Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) * = den Tieren dauernd zugängliche Liege- + Fress- + Lauffläche inner- und ausserhalb des Stalls", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_2)
	cpA6_3 := models.ControlPoint{ControlPointId: "03", Abbreviation: "Alle übrigen Laufhöfe: Gesamte und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Laufhof-Skizze, die für die Berechnung der gesamten und der ungedeckten Laufhoffläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte gesamte und ungedeckte Laufhoffläche wurden korrekt berechnet (= nachrechnen) Kategorien A und B: Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) Kategorien C: mind. 25 % der Laufhoffläche müssen ungedeckt sein Kategorien D: mind. 50 % der Laufhoffläche müssen ungedeckt sein", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_3)
	cpA6_4 := models.ControlPoint{ControlPointId: "04", Abbreviation: "Die Laufhof-Skizze entspricht den Anforderungen, ist verifiziert und aktuell", ControlPoint: "Anforderungen: siehe 02 Aktuell: grober visueller Abgleich", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_4)
	cpA6_5 := models.ControlPoint{ControlPointId: "05", Abbreviation: "Aktuelle Tierzahl pro Auslauf kleiner oder gleich max. Tierzahl auf aktueller und verifizierter Laufhof-Skizze", ControlPoint: "Anzahl Tiere, denen aktuell gleichzeitig Auslauf gewährt wird: Tiere Zählen bzw. Selbstdeklaration des Landwirts verifizieren", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_5)
	cpA6_6 := models.ControlPoint{ControlPointId: "06", Abbreviation: "1.11.-28.2.: kein Schattennetz", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_6)
	cpA6_7 := models.ControlPoint{ControlPointId: "07", Abbreviation: "Unbefestigte Laufhöfe: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei unbefestigten Laufhöfen: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_7)
	cpA6_8 := models.ControlPoint{ControlPointId: "08", Abbreviation: "Liegebereich ohne Perforierung und ausreichend mit geeigneter Einstreue bedeckt", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_8)
	cpA6_11 := models.ControlPoint{ControlPointId: "11", Abbreviation: "Weiden: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei Weiden: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_11)
	cpA6_12 := models.ControlPoint{ControlPointId: "12", Abbreviation: "Weide kann an Weidetagen ca. 25 % des TS-Verzehrs decken", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_12)
	cpA6_141 := models.ControlPoint{ControlPointId: "14.1", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_141)
	cpA6_1411 := models.ControlPoint{ControlPointId: "14.1.1", Abbreviation: "Auslauf-Dokumentation entspricht den Anforderungen", ControlPoint: "Anforderungen an die Dokumentation: 1. Auslauf je Auslaufgruppe bzw. je Tier eingetragen 2. a) Auslauf nach spätestens 3 Tagen eingetragen bzw. b) Anfang und Ende von Zeitspannen eingetragen, während denen die Tiere 24 h am Tag Zugang haben zu - einer Weide (01.05.-31.10.) bzw. - einem Laufhof (01.11.-30.04.) Ist die Einhaltung der Auslaufbestimmungen während des ganzen Jahres durch das Haltungssystem gewährleistet, ist kein Journal erforderlich", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_1411)
	cpA6_1412 := models.ControlPoint{ControlPointId: "14.1.2", Abbreviation: "genügend Weide- bzw. bis Auslauftage",
		ControlPoint: "Allen Tieren der Kategorie wurde vom 01.05 bis zum 31.10. an mindestens 26 Tagen pro Monat und vom 01.11. bis zum 30.04. an mindestens 13 Tagen pro Monat Auslauf auf einer Weide gewährt oder zulässige Abweichungen gemäss DZV", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_1412)
	cpA6_142 := models.ControlPoint{ControlPointId: "14.2", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_142)
	cpA6_1421 := models.ControlPoint{ControlPointId: "14.2.1", Abbreviation: "Alternativ-Variante für betreffende Tiere zulässig und Laufhof dauernd zugänglich", ControlPoint: "Auslauf-Alternativ-Variante für Tiere der Rindergattung, - die gemästet werden; - männliche Zuchttiere; und - bis 160 Tage alte weibliche Zuchttiere Alle Tiere d. Kategorie hatten während des ganzen Jahres dauernd (24 h pro Tag) Zugang zu Laufhof oder zulässige Abweichungen gemäss DZV", PointGroup: &pg6}
	services.CreateControlPoint(&cpA6_1421)

	services.CreateLack(&models.Lack{Name: "Laufhof befindet sich nicht im Freien", Points: 110, ControlPoint: &cpA6_1})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA6_2})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA6_2})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA6_3})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA6_3})

	services.CreateLack(&models.Lack{Name: "Laufhof-Skizze kann nicht vorgewiesen werden, entspricht nicht den Anforderungen, ist nicht verifiziert oder nicht aktuell", Francs: 200, ControlPoint: &cpA6_4})

	services.CreateLack(&models.Lack{Name: "Aktuelle Tierzahl pro Auslauf grösser als maximale Tierzahl auf aktueller und verifizierter Laufhof-Skizze", Points: 110, ControlPoint: &cpA6_5})

	services.CreateLack(&models.Lack{Name: "Vom 1.11 ' 28.2. ein Schattennetz", Points: 10, ControlPoint: &cpA6_6})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA6_7})

	services.CreateLack(&models.Lack{Name: "Zu wenig geeignete Einstreu", Points: 10, ControlPoint: &cpA6_8})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig geeignete Einstreu", Points: 40, ControlPoint: &cpA6_8})
	services.CreateLack(&models.Lack{Name: "Keine geeignete Einstreu", Points: 110, ControlPoint: &cpA6_8})
	services.CreateLack(&models.Lack{Name: "Liegebereich(e) mit Perforation", Points: 110, ControlPoint: &cpA6_8})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA6_11})

	services.CreateLack(&models.Lack{Name: "Weide kann an Weidetagen nicht ca. 25% des TS-Verzehrs decken", Points: 60, ControlPoint: &cpA6_12})

	services.CreateLack(&models.Lack{Name: "Auslauf-Dokumentation entspricht nicht den Anforderungen", Francs: 200, ControlPoint: &cpA6_1411})

	services.CreateLack(&models.Lack{Name: "zu wenig Tage mit Zugang zur Weide- bzw. zum Laufhof nachgewiesen", Points: 4, Computed:true, ControlPoint: &cpA6_1412})

	services.CreateLack(&models.Lack{Name: "Alternativ Variante für betreffende Tiere nicht zulässig oder Laufhof nicht dauernd zugänglich", Points: 110, ControlPoint: &cpA6_1421})

	// A7 - Rinder - männliche Tiere, über 365- 730 Tage alt
	pg7 := models.PointGroup{PointGroupId: "A7", Abbreviation: "Rinder - männliche Tiere, über 365- 730 Tage alt", PointGroup: "männliche Tiere, über 365 bis 730 Tage alt", PointGroupCode: 1129, ControlCategory: &cc}
	services.CreatePointGroup(&pg7)

	cpA7_1 := models.ControlPoint{ControlPointId: "01", Abbreviation: "Laufhof befindet sich im Freien", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_1)
	cpA7_2 := models.ControlPoint{ControlPointId: "02", Abbreviation: "Laufhöfe, die für das Rindvieh dauernd zugänglich sind: Gesamtfläche und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Stall-Laufhof-Skizze, die für die Berechnung der Gesamtfläche* und der ungedeckten Fläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte Gesamtfläche* und ungedeckte Fläche wurden korrekt berechnet (= nachrechnen) Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) * = den Tieren dauernd zugängliche Liege- + Fress- + Lauffläche inner- und ausserhalb des Stalls", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_2)
	cpA7_3 := models.ControlPoint{ControlPointId: "03", Abbreviation: "Alle übrigen Laufhöfe: Gesamte und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Laufhof-Skizze, die für die Berechnung der gesamten und der ungedeckten Laufhoffläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte gesamte und ungedeckte Laufhoffläche wurden korrekt berechnet (= nachrechnen) Kategorien A und B: Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) Kategorien C: mind. 25 % der Laufhoffläche müssen ungedeckt sein Kategorien D: mind. 50 % der Laufhoffläche müssen ungedeckt sein", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_3)
	cpA7_4 := models.ControlPoint{ControlPointId: "04", Abbreviation: "Die Laufhof-Skizze entspricht den Anforderungen, ist verifiziert und aktuell", ControlPoint: "Anforderungen: siehe 02 Aktuell: grober visueller Abgleich", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_4)
	cpA7_5 := models.ControlPoint{ControlPointId: "05", Abbreviation: "Aktuelle Tierzahl pro Auslauf kleiner oder gleich max. Tierzahl auf aktueller und verifizierter Laufhof-Skizze", ControlPoint: "Anzahl Tiere, denen aktuell gleichzeitig Auslauf gewährt wird: Tiere Zählen bzw. Selbstdeklaration des Landwirts verifizieren", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_5)
	cpA7_6 := models.ControlPoint{ControlPointId: "06", Abbreviation: "1.11.-28.2.: kein Schattennetz", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_6)
	cpA7_7 := models.ControlPoint{ControlPointId: "07", Abbreviation: "Unbefestigte Laufhöfe: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei unbefestigten Laufhöfen: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_7)
	cpA7_8 := models.ControlPoint{ControlPointId: "08", Abbreviation: "Liegebereich ohne Perforierung und ausreichend mit geeigneter Einstreue bedeckt", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_8)
	cpA7_11 := models.ControlPoint{ControlPointId: "11", Abbreviation: "Weiden: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei Weiden: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_11)
	cpA7_12 := models.ControlPoint{ControlPointId: "12", Abbreviation: "Weide kann an Weidetagen ca. 25 % des TS-Verzehrs decken", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_12)
	cpA7_141 := models.ControlPoint{ControlPointId: "14.1", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_141)
	cpA7_1411 := models.ControlPoint{ControlPointId: "14.1.1", Abbreviation: "Auslauf-Dokumentation entspricht den Anforderungen", ControlPoint: "Anforderungen an die Dokumentation: 1. Auslauf je Auslaufgruppe bzw. je Tier eingetragen 2. a) Auslauf nach spätestens 3 Tagen eingetragen bzw. b) Anfang und Ende von Zeitspannen eingetragen, während denen die Tiere 24 h am Tag Zugang haben zu - einer Weide (01.05.-31.10.) bzw. - einem Laufhof (01.11.-30.04.) Ist die Einhaltung der Auslaufbestimmungen während des ganzen Jahres durch das Haltungssystem gewährleistet, ist kein Journal erforderlich", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_1411)
	cpA7_1412 := models.ControlPoint{ControlPointId: "14.1.2", Abbreviation: "genügend Weide- bzw. bis Auslauftage",
		ControlPoint: "Allen Tieren der Kategorie wurde vom 01.05 bis zum 31.10. an mindestens 26 Tagen pro Monat und vom 01.11. bis zum 30.04. an mindestens 13 Tagen pro Monat Auslauf auf einer Weide gewährt oder zulässige Abweichungen gemäss DZV", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_1412)
	cpA7_142 := models.ControlPoint{ControlPointId: "14.2", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_142)
	cpA7_1421 := models.ControlPoint{ControlPointId: "14.2.1", Abbreviation: "Alternativ-Variante für betreffende Tiere zulässig und Laufhof dauernd zugänglich", ControlPoint: "Auslauf-Alternativ-Variante für Tiere der Rindergattung, - die gemästet werden; - männliche Zuchttiere; und - bis 160 Tage alte weibliche Zuchttiere Alle Tiere d. Kategorie hatten während des ganzen Jahres dauernd (24 h pro Tag) Zugang zu Laufhof oder zulässige Abweichungen gemäss DZV", PointGroup: &pg7}
	services.CreateControlPoint(&cpA7_1421)

	services.CreateLack(&models.Lack{Name: "Laufhof befindet sich nicht im Freien", Points: 110, ControlPoint: &cpA7_1})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA7_2})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA7_2})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA7_3})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA7_3})

	services.CreateLack(&models.Lack{Name: "Laufhof-Skizze kann nicht vorgewiesen werden, entspricht nicht den Anforderungen, ist nicht verifiziert oder nicht aktuell", Francs: 200, ControlPoint: &cpA7_4})

	services.CreateLack(&models.Lack{Name: "Aktuelle Tierzahl pro Auslauf grösser als maximale Tierzahl auf aktueller und verifizierter Laufhof-Skizze", Points: 110, ControlPoint: &cpA7_5})

	services.CreateLack(&models.Lack{Name: "Vom 1.11 ' 28.2. ein Schattennetz", Points: 10, ControlPoint: &cpA7_6})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA7_7})

	services.CreateLack(&models.Lack{Name: "Zu wenig geeignete Einstreu", Points: 10, ControlPoint: &cpA7_8})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig geeignete Einstreu", Points: 40, ControlPoint: &cpA7_8})
	services.CreateLack(&models.Lack{Name: "Keine geeignete Einstreu", Points: 110, ControlPoint: &cpA7_8})
	services.CreateLack(&models.Lack{Name: "Liegebereich(e) mit Perforation", Points: 110, ControlPoint: &cpA7_8})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA7_11})

	services.CreateLack(&models.Lack{Name: "Weide kann an Weidetagen nicht ca. 25% des TS-Verzehrs decken", Points: 60, ControlPoint: &cpA7_12})

	services.CreateLack(&models.Lack{Name: "Auslauf-Dokumentation entspricht nicht den Anforderungen", Francs: 200, ControlPoint: &cpA7_1411})

	services.CreateLack(&models.Lack{Name: "zu wenig Tage mit Zugang zur Weide- bzw. zum Laufhof nachgewiesen", Points: 4, Computed:true, ControlPoint: &cpA7_1412})

	services.CreateLack(&models.Lack{Name: "Alternativ Variante für betreffende Tiere nicht zulässig oder Laufhof nicht dauernd zugänglich", Points: 110, ControlPoint: &cpA7_1421})

	// A8 - Rinder - männliche Tiere, über 160-365 Tage alt
	pg8 := models.PointGroup{PointGroupId: "A8", Abbreviation: "Rinder - männliche Tiere, über 160-365 Tage alt", PointGroup: "männliche Tiere, über 160 bis 365 Tage alt", PointGroupCode: 1143, ControlCategory: &cc}
	services.CreatePointGroup(&pg8)

	cpA8_1 := models.ControlPoint{ControlPointId: "01", Abbreviation: "Laufhof befindet sich im Freien", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_1)
	cpA8_2 := models.ControlPoint{ControlPointId: "02", Abbreviation: "Laufhöfe, die für das Rindvieh dauernd zugänglich sind: Gesamtfläche und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Stall-Laufhof-Skizze, die für die Berechnung der Gesamtfläche* und der ungedeckten Fläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte Gesamtfläche* und ungedeckte Fläche wurden korrekt berechnet (= nachrechnen) Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) * = den Tieren dauernd zugängliche Liege- + Fress- + Lauffläche inner- und ausserhalb des Stalls", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_2)
	cpA8_3 := models.ControlPoint{ControlPointId: "03", Abbreviation: "Alle übrigen Laufhöfe: Gesamte und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Laufhof-Skizze, die für die Berechnung der gesamten und der ungedeckten Laufhoffläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte gesamte und ungedeckte Laufhoffläche wurden korrekt berechnet (= nachrechnen) Kategorien A und B: Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) Kategorien C: mind. 25 % der Laufhoffläche müssen ungedeckt sein Kategorien D: mind. 50 % der Laufhoffläche müssen ungedeckt sein", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_3)
	cpA8_4 := models.ControlPoint{ControlPointId: "04", Abbreviation: "Die Laufhof-Skizze entspricht den Anforderungen, ist verifiziert und aktuell", ControlPoint: "Anforderungen: siehe 02 Aktuell: grober visueller Abgleich", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_4)
	cpA8_5 := models.ControlPoint{ControlPointId: "05", Abbreviation: "Aktuelle Tierzahl pro Auslauf kleiner oder gleich max. Tierzahl auf aktueller und verifizierter Laufhof-Skizze", ControlPoint: "Anzahl Tiere, denen aktuell gleichzeitig Auslauf gewährt wird: Tiere Zählen bzw. Selbstdeklaration des Landwirts verifizieren", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_5)
	cpA8_6 := models.ControlPoint{ControlPointId: "06", Abbreviation: "1.11.-28.2.: kein Schattennetz", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_6)
	cpA8_7 := models.ControlPoint{ControlPointId: "07", Abbreviation: "Unbefestigte Laufhöfe: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei unbefestigten Laufhöfen: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_7)
	cpA8_8 := models.ControlPoint{ControlPointId: "08", Abbreviation: "Liegebereich ohne Perforierung und ausreichend mit geeigneter Einstreue bedeckt", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_8)
	cpA8_11 := models.ControlPoint{ControlPointId: "11", Abbreviation: "Weiden: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei Weiden: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_11)
	cpA8_12 := models.ControlPoint{ControlPointId: "12", Abbreviation: "Weide kann an Weidetagen ca. 25 % des TS-Verzehrs decken", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_12)
	cpA8_141 := models.ControlPoint{ControlPointId: "14.1", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_141)
	cpA8_1411 := models.ControlPoint{ControlPointId: "14.1.1", Abbreviation: "Auslauf-Dokumentation entspricht den Anforderungen", ControlPoint: "Anforderungen an die Dokumentation: 1. Auslauf je Auslaufgruppe bzw. je Tier eingetragen 2. a) Auslauf nach spätestens 3 Tagen eingetragen bzw. b) Anfang und Ende von Zeitspannen eingetragen, während denen die Tiere 24 h am Tag Zugang haben zu - einer Weide (01.05.-31.10.) bzw. - einem Laufhof (01.11.-30.04.) Ist die Einhaltung der Auslaufbestimmungen während des ganzen Jahres durch das Haltungssystem gewährleistet, ist kein Journal erforderlich", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_1411)
	cpA8_1412 := models.ControlPoint{ControlPointId: "14.1.2", Abbreviation: "genügend Weide- bzw. bis Auslauftage",
		ControlPoint: "Allen Tieren der Kategorie wurde vom 01.05 bis zum 31.10. an mindestens 26 Tagen pro Monat und vom 01.11. bis zum 30.04. an mindestens 13 Tagen pro Monat Auslauf auf einer Weide gewährt oder zulässige Abweichungen gemäss DZV", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_1412)
	cpA8_142 := models.ControlPoint{ControlPointId: "14.2", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_142)
	cpA8_1421 := models.ControlPoint{ControlPointId: "14.2.1", Abbreviation: "Alternativ-Variante für betreffende Tiere zulässig und Laufhof dauernd zugänglich", ControlPoint: "Auslauf-Alternativ-Variante für Tiere der Rindergattung, - die gemästet werden; - männliche Zuchttiere; und - bis 160 Tage alte weibliche Zuchttiere Alle Tiere d. Kategorie hatten während des ganzen Jahres dauernd (24 h pro Tag) Zugang zu Laufhof oder zulässige Abweichungen gemäss DZV", PointGroup: &pg8}
	services.CreateControlPoint(&cpA8_1421)

	services.CreateLack(&models.Lack{Name: "Laufhof befindet sich nicht im Freien", Points: 110, ControlPoint: &cpA8_1})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA8_2})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA8_2})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA8_3})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA8_3})

	services.CreateLack(&models.Lack{Name: "Laufhof-Skizze kann nicht vorgewiesen werden, entspricht nicht den Anforderungen, ist nicht verifiziert oder nicht aktuell", Francs: 200, ControlPoint: &cpA8_4})

	services.CreateLack(&models.Lack{Name: "Aktuelle Tierzahl pro Auslauf grösser als maximale Tierzahl auf aktueller und verifizierter Laufhof-Skizze", Points: 110, ControlPoint: &cpA8_5})

	services.CreateLack(&models.Lack{Name: "Vom 1.11 ' 28.2. ein Schattennetz", Points: 10, ControlPoint: &cpA8_6})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA8_7})

	services.CreateLack(&models.Lack{Name: "Zu wenig geeignete Einstreu", Points: 10, ControlPoint: &cpA8_8})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig geeignete Einstreu", Points: 40, ControlPoint: &cpA8_8})
	services.CreateLack(&models.Lack{Name: "Keine geeignete Einstreu", Points: 110, ControlPoint: &cpA8_8})
	services.CreateLack(&models.Lack{Name: "Liegebereich(e) mit Perforation", Points: 110, ControlPoint: &cpA8_8})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA8_11})

	services.CreateLack(&models.Lack{Name: "Weide kann an Weidetagen nicht ca. 25% des TS-Verzehrs decken", Points: 60, ControlPoint: &cpA8_12})

	services.CreateLack(&models.Lack{Name: "Auslauf-Dokumentation entspricht nicht den Anforderungen", Francs: 200, ControlPoint: &cpA8_1411})

	services.CreateLack(&models.Lack{Name: "zu wenig Tage mit Zugang zur Weide- bzw. zum Laufhof nachgewiesen", Points: 4, Computed:true, ControlPoint: &cpA8_1412})

	services.CreateLack(&models.Lack{Name: "Alternativ Variante für betreffende Tiere nicht zulässig oder Laufhof nicht dauernd zugänglich", Points: 110, ControlPoint: &cpA8_1421})

	// A9 - Rinder - männliche Tiere, bis 160 Tage alt
	pg9 := models.PointGroup{PointGroupId: "A8", Abbreviation: "Rinder - männliche Tiere, bis 160 Tage alt", PointGroup: "männliche Tiere, bis 160 Tage alt", PointGroupCode: 1144, ControlCategory: &cc}
	services.CreatePointGroup(&pg9)

	cpA9_1 := models.ControlPoint{ControlPointId: "01", Abbreviation: "Laufhof befindet sich im Freien", PointGroup: &pg9}
	services.CreateControlPoint(&cpA9_1)
	cpA9_2 := models.ControlPoint{ControlPointId: "02", Abbreviation: "Laufhöfe, die für das Rindvieh dauernd zugänglich sind: Gesamtfläche und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Stall-Laufhof-Skizze, die für die Berechnung der Gesamtfläche* und der ungedeckten Fläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte Gesamtfläche* und ungedeckte Fläche wurden korrekt berechnet (= nachrechnen) Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) * = den Tieren dauernd zugängliche Liege- + Fress- + Lauffläche inner- und ausserhalb des Stalls", PointGroup: &pg9}
	services.CreateControlPoint(&cpA9_2)
	cpA9_3 := models.ControlPoint{ControlPointId: "03", Abbreviation: "Alle übrigen Laufhöfe: Gesamte und ungedeckte Laufhoffläche - wurden korrekt vermessen und berechnet und - entsprechen den Anforderungen", ControlPoint: "Alle Masse auf der aktuellen Laufhof-Skizze, die für die Berechnung der gesamten und der ungedeckten Laufhoffläche relevant sind, wurden korrekt vermessen (= nachmessen) Die auf der Skizze vermerkte gesamte und ungedeckte Laufhoffläche wurden korrekt berechnet (= nachrechnen) Kategorien A und B: Für die auf der Skizze vermerkte maximale Tierzahl sind die Anforderungen erfüllt (= nachrechnen) Kategorien C: mind. 25 % der Laufhoffläche müssen ungedeckt sein Kategorien D: mind. 50 % der Laufhoffläche müssen ungedeckt sein", PointGroup: &pg9}
	services.CreateControlPoint(&cpA9_3)
	cpA9_4 := models.ControlPoint{ControlPointId: "04", Abbreviation: "Die Laufhof-Skizze entspricht den Anforderungen, ist verifiziert und aktuell", ControlPoint: "Anforderungen: siehe 02 Aktuell: grober visueller Abgleich", PointGroup: &pg9}
	services.CreateControlPoint(&cpA9_4)
	cpA9_5 := models.ControlPoint{ControlPointId: "05", Abbreviation: "Aktuelle Tierzahl pro Auslauf kleiner oder gleich max. Tierzahl auf aktueller und verifizierter Laufhof-Skizze", ControlPoint: "Anzahl Tiere, denen aktuell gleichzeitig Auslauf gewährt wird: Tiere Zählen bzw. Selbstdeklaration des Landwirts verifizieren", PointGroup: &pg9}
	services.CreateControlPoint(&cpA9_5)
	cpA9_6 := models.ControlPoint{ControlPointId: "06", Abbreviation: "1.11.-28.2.: kein Schattennetz", PointGroup: &pg9}
	services.CreateControlPoint(&cpA9_6)
	cpA9_7 := models.ControlPoint{ControlPointId: "07", Abbreviation: "Unbefestigte Laufhöfe: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei unbefestigten Laufhöfen: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg9}
	services.CreateControlPoint(&cpA9_7)
	cpA9_8 := models.ControlPoint{ControlPointId: "08", Abbreviation: "Liegebereich ohne Perforierung und ausreichend mit geeigneter Einstreue bedeckt", PointGroup: &pg9}
	services.CreateControlPoint(&cpA9_8)
	cpA9_9 := models.ControlPoint{ControlPointId: "08", Abbreviation: "Bis 160 Tage alte Tiere nicht fixiert", PointGroup: &pg9}
	services.CreateControlPoint(&cpA9_9)
	cpA9_11 := models.ControlPoint{ControlPointId: "11", Abbreviation: "Weiden: alle morastigen Stellen ausgezäunt", ControlPoint: "Ausnahmen bei Weiden: z.B. Suhle für Wasserbüffel mit schriftlicher Bewilligung des Kantons (z.B. des Amts für Umwelt)", PointGroup: &pg9}
	services.CreateControlPoint(&cpA9_11)
	cpA9_12 := models.ControlPoint{ControlPointId: "12", Abbreviation: "Weide kann an Weidetagen ca. 25 % des TS-Verzehrs decken", PointGroup: &pg9}
	services.CreateControlPoint(&cpA9_12)
	cpA9_141 := models.ControlPoint{ControlPointId: "14.1", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg9}
	services.CreateControlPoint(&cpA9_141)
	cpA9_1411 := models.ControlPoint{ControlPointId: "14.1.1", Abbreviation: "Auslauf-Dokumentation entspricht den Anforderungen", ControlPoint: "Anforderungen an die Dokumentation: 1. Auslauf je Auslaufgruppe bzw. je Tier eingetragen 2. a) Auslauf nach spätestens 3 Tagen eingetragen bzw. b) Anfang und Ende von Zeitspannen eingetragen, während denen die Tiere 24 h am Tag Zugang haben zu - einer Weide (01.05.-31.10.) bzw. - einem Laufhof (01.11.-30.04.) Ist die Einhaltung der Auslaufbestimmungen während des ganzen Jahres durch das Haltungssystem gewährleistet, ist kein Journal erforderlich", PointGroup: &pg9}
	services.CreateControlPoint(&cpA9_1411)
	cpA9_1412 := models.ControlPoint{ControlPointId: "14.1.2", Abbreviation: "genügend Weide- bzw. bis Auslauftage",
		ControlPoint: "Allen Tieren der Kategorie wurde vom 01.05 bis zum 31.10. an mindestens 26 Tagen pro Monat und vom 01.11. bis zum 30.04. an mindestens 13 Tagen pro Monat Auslauf auf einer Weide gewährt oder zulässige Abweichungen gemäss DZV", PointGroup: &pg9}
	services.CreateControlPoint(&cpA9_1412)

	cpA9_142 := models.ControlPoint{ControlPointId: "14.2", Abbreviation: "Auslauf: Alternativ-Variante (ohne Weide)", PointGroup: &pg9}
	services.CreateControlPoint(&cpA9_142)
	cpA9_1421 := models.ControlPoint{ControlPointId: "14.2.1", Abbreviation: "Alternativ-Variante für betreffende Tiere zulässig und Laufhof dauernd zugänglich", ControlPoint: "Auslauf-Alternativ-Variante für Tiere der Rindergattung, - die gemästet werden; - männliche Zuchttiere; und - bis 160 Tage alte weibliche Zuchttiere Alle Tiere d. Kategorie hatten während des ganzen Jahres dauernd (24 h pro Tag) Zugang zu Laufhof oder zulässige Abweichungen gemäss DZV", PointGroup: &pg9}
	services.CreateControlPoint(&cpA9_1421)

	services.CreateLack(&models.Lack{Name: "Laufhof befindet sich nicht im Freien", Points: 110, ControlPoint: &cpA9_1})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA9_2})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA9_2})

	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche umweniger als 10 %", Points: 60, ControlPoint: &cpA9_3})
	services.CreateLack(&models.Lack{Name: "Nachgemessene Gesamtfläche oder ungedeckte Laufhoffläche unterschreitet Mindestfläche um10 oder mehr %", Points: 110, ControlPoint: &cpA9_3})

	services.CreateLack(&models.Lack{Name: "Laufhof-Skizze kann nicht vorgewiesen werden, entspricht nicht den Anforderungen, ist nicht verifiziert oder nicht aktuell", Francs: 200, ControlPoint: &cpA9_4})

	services.CreateLack(&models.Lack{Name: "Aktuelle Tierzahl pro Auslauf grösser als maximale Tierzahl auf aktueller und verifizierter Laufhof-Skizze", Points: 110, ControlPoint: &cpA9_5})

	services.CreateLack(&models.Lack{Name: "Vom 1.11 ' 28.2. ein Schattennetz", Points: 10, ControlPoint: &cpA9_6})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA9_7})

	services.CreateLack(&models.Lack{Name: "Zu wenig geeignete Einstreu", Points: 10, ControlPoint: &cpA9_8})
	services.CreateLack(&models.Lack{Name: "Viel zu wenig geeignete Einstreu", Points: 40, ControlPoint: &cpA9_8})
	services.CreateLack(&models.Lack{Name: "Keine geeignete Einstreu", Points: 110, ControlPoint: &cpA9_8})
	services.CreateLack(&models.Lack{Name: "Liegebereich(e) mit Perforation", Points: 110, ControlPoint: &cpA9_8})

	services.CreateLack(&models.Lack{Name: "Alter des fixierten Tiers weniger als 4 Monate", Points: 110, ControlPoint: &cpA9_9})
	services.CreateLack(&models.Lack{Name: "2015 und 2016: Alter des fixierten Tiers älter als 4 Monate", ControlPoint: &cpA9_9})

	services.CreateLack(&models.Lack{Name: "morastige Stelle(n) nicht ausgezäunt", Points: 10, ControlPoint: &cpA9_11})

	services.CreateLack(&models.Lack{Name: "Weide kann an Weidetagen nicht ca. 25% des TS-Verzehrs decken", Points: 60, ControlPoint: &cpA9_12})

	services.CreateLack(&models.Lack{Name: "Auslauf-Dokumentation entspricht nicht den Anforderungen", Francs: 200, ControlPoint: &cpA9_1411})

	services.CreateLack(&models.Lack{Name: "zu wenig Tage mit Zugang zur Weide- bzw. zum Laufhof nachgewiesen", Points: 4, Computed:true, ControlPoint: &cpA9_1412})

	services.CreateLack(&models.Lack{Name: "Alternativ Variante für betreffende Tiere nicht zulässig oder Laufhof nicht dauernd zugänglich", Points: 110, ControlPoint: &cpA9_1421})
}

func Seed_Users() {
	if cnt, _ := services.CountRoles(); cnt > 0 {
		return
	}
	admin := models.Role{Name: "Admin"}
	farmer := models.Role{Name: "Farmer"}
	canton := models.Role{Name: "Canton"}
	inspector := models.Role{Name: "Inspector"}

	services.CreateRole(&admin)
	services.CreateRole(&farmer)
	services.CreateRole(&canton)
	services.CreateRole(&inspector)

	if cnt, _ := services.CountUsers(); cnt > 0 {
		return
	}

	roles := make([]*models.Role, 1)
	roles[0] = &farmer
	services.CreateUser(&models.User{Username: "farmer1", Password: "farmer1", TVD: 1015010, Roles: roles, Email: "farmer1@apayment.ch"})
	services.CreateUser(&models.User{Username: "farmer2", Password: "farmer2", TVD: 1015010, Roles: roles, Email: "farmer2@apayment.ch"})
	services.CreateUser(&models.User{Username: "field", Password: "field", TVD: 0, Lastname: "Field", Firstname: "Test", Roles: roles, Email: "field@apayment.ch"})

	roles = make([]*models.Role, 1)
	roles[0] = &inspector
	services.CreateUser(&models.User{Username: "inspect", Password: "inspect", Firstname: "Inspector", Lastname: "Gadget", Roles: roles, Email: "inspect@apayment.ch"})

	roles = make([]*models.Role, 1)
	roles[0] = &admin
	services.CreateUser(&models.User{Username: "admin", Password: "admin", Firstname: "Admin", Lastname: "Admin", Roles: roles, Email: "admin@apayment.ch"})

	roles = make([]*models.Role, 1)
	roles[0] = &canton
	services.CreateUser(&models.User{Username: "canton", Password: "canton", Firstname: "Maria", Lastname: "Walder", Roles: roles, Email: "canton@apayment.ch"})

}

func Seed_LegalForm() {
	if cnt, _ := services.CountLegalForms(); cnt > 0 {
		return
	}
	services.CreateLegalForm(&models.LegalForm{Code: 1, Name: "Natürtliche Person"})
	services.CreateLegalForm(&models.LegalForm{Code: 2, Name: "Einfache Gesellschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 3, Name: "Kollektivgesellschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 4, Name: "Kommanditgesellschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 5, Name: "Kommanditaktiengesellschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 6, Name: "Aktiengesellschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 7, Name: "Gesellschaft mit beschränkter Haftung"})
	services.CreateLegalForm(&models.LegalForm{Code: 8, Name: "Genossenschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 9, Name: "Verein, Vereinigung"})
	services.CreateLegalForm(&models.LegalForm{Code: 10, Name: "Stiftung"})

	services.CreateLegalForm(&models.LegalForm{Code: 24, Name: "Öffentlich-rechtliche Körperschaft(Verwaltung)"})
	services.CreateLegalForm(&models.LegalForm{Code: 25, Name: "Staatlich anerkannte Landeskirche"})
	services.CreateLegalForm(&models.LegalForm{Code: 30, Name: "Bund (Betrieb)"})
	services.CreateLegalForm(&models.LegalForm{Code: 31, Name: "Kanton (Betrieb)"})
	services.CreateLegalForm(&models.LegalForm{Code: 32, Name: "Bezirk (Betrieb)"})

	services.CreateLegalForm(&models.LegalForm{Code: 33, Name: "Gemeinde (Betrieb)"})
	services.CreateLegalForm(&models.LegalForm{Code: 34, Name: "Öffentlich-rechtliche Körperschaft (Betrieb)"})
	services.CreateLegalForm(&models.LegalForm{Code: 99, Name: "Nicht zugeteilt"})
}

func Seed_PlantType() {
	if cnt, _ := services.CountPlantTypes(); cnt > 0 {
		return
	}
	services.CreatePlantType(&models.PlantType{Code: 1, Name: "Ganzjahresbetrieb"})
	services.CreatePlantType(&models.PlantType{Code: 2, Name: "Produktsstätte"})
	services.CreatePlantType(&models.PlantType{Code: 3, Name: "Betriebsgemeinschaft"})
}

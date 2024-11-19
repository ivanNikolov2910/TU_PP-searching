# Тема М1-1. Сравнение на производителността на алгоритми за паралелно търсене 

## Описание на проекта: 
    В този проект ще реализирате и ще сравните производителността на различни алгоритми за търсене, използвайки различни техники за паралелизация. Целта е да се анализира как се представят различните алгоритми за паралелно търсене върху различни набори от данни и хардуерни конфигурации. 

## Компоненти на проекта:

### Избор на алгоритъм: 
    Реализирайте няколко алгоритъма за паралелно търсене. Разгледайте алгоритми като паралелно линейно търсене, паралелно двоично търсене, паралелно търсене на базата на хеш-таблици и паралелно търсене на базата на дървовидни структури.
    * parallel linear search: https://www.geeksforgeeks.org/linear-search-using-multi-threading/
    * parallel binary search: https://codeforces.com/blog/entry/45578
    * parallel hash-based search: https://library.fiveable.me/key-terms/exascale-computing/parallel-hash-based-search
    * parallel tree search: https://github.com/mtz1024/parallel-tree-search
### Техники за паралелизация: 
    Изследвайте различни техники за паралелизация за CPU и/или GPU.
### Бенчмаркинг: 
    Разработете фреймуърк за сравнителен анализ, която ви позволява да измервате времето за изпълнение, мащабируемостта и ефективността на всеки алгоритъм за търсене. Този фреймуърк трябва да генерира случайни входни данни с различна размерност и характеристики (напр. сортирани, обратно сортирани, случайни последователности) за тестване.
### Хардуерна конфигурация: 
    Тествайте алгоритмите за търсене на различни хардуерни конфигурации, като например многоядрен процесор и графичен процесор (ако има такъв!).
### Показатели за производителност: 
    Определете показатели за производителност за целите на провеждания сравнителен анализ, включително време за изпълнение, ускорение и ефективност. Сравнете производителността на всеки алгоритъм при различни размери на входните набори от данни.
### Потребителски интерфейс: 
    Разработване на удобен за потребителя интерфейс, който позволява на потребителите да зареждат входни данни, да избират паралелни алгоритми и да сравняват показателите им за производителност.
### Анализ: 
    Анализирайте резултатите и направете изводи за това кои алгоритми за паралелно търсене са най-подходящи за различните сценарии. Вземете предвид фактори като размер на входните данни, разпределение на данните и хардуерни ресурси.
### Оптимизации: 
    Разгледайте потенциалните възможности за оптимизация за алгоритмите за търсене, които са се представили по-неефективно в сравнение с останалите. Експериментирайте с различни стратегии за паралелизация или алгоритмични подобрения, за да подобрите тяхната производителност.
### Документация и презентиране: 
    Изгответе документация, която да обхваща всички изброени компоненти на проекта, да съдържа описание на изследваните алгоритми и стратегии за тяхната паралелизация, експериментални резултати и резултати от анализи. Включете визуализации, графики и таблици, за илюстрация на резултатите. Създайте презентация за представяне на резултатите от проекта.